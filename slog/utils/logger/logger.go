package logger

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

// Logger: ログ出力機能を提供する構造体
type Logger struct {
	*slog.Logger
}

// SystemInfo: システム情報を保持する構造体
type SystemInfo struct {
	Environment string `json:"environment"`
	Service     string `json:"service"`
	Hostname    string `json:"hostname"`
}

// NewSystemInfo: SystemInfoの新しいインスタンスを作成
func NewSystemInfo(env string) SystemInfo {
	hostname, _ := os.Hostname()
	return SystemInfo{
		Environment: env,
		Service:     fmt.Sprintf("%s-slog-server", env),
		Hostname:    hostname,
	}
}

// HTTPRequestInfo: HTTPリクエスト情報を保持する構造体
type HTTPRequestInfo struct {
	Method     string `json:"method"`
	Path       string `json:"path"`
	Status     int    `json:"status"`
	Latency    string `json:"latency"`
	UserAgent  string `json:"user_agent"`
	Referer    string `json:"referer"`
	RemoteAddr string `json:"remote_addr"`
	RequestID  string `json:"request_id"`
}

// NewInitialHTTPRequestInfo: HTTPRequestInfoの新しいインスタンスを作成
func NewInitialHTTPRequestInfo(r *http.Request, start time.Time, statusCode int) HTTPRequestInfo {
	// 処理時間の計算
	duration := time.Since(start)

	// コンテキストからrequest_idを取得
	requestID, ok := GetRequestID(r.Context())
	if !ok {
		// フォールバック: タイムスタンプベースのID
		requestID = uuid.New().String()
	}

	return HTTPRequestInfo{
		Method:     r.Method,
		Path:       r.URL.Path,
		Status:     statusCode,
		Latency:    duration.String(),
		UserAgent:  r.UserAgent(),
		Referer:    r.Referer(),
		RemoteAddr: r.RemoteAddr,
		RequestID:  requestID,
	}
}

// AuthorizedInfo: 認可後に得られる情報を保持する構造体
type AuthorizedInfo struct {
	Role     string `json:"role"`
	TenantID string `json:"tenant_id"`
	MemberID string `json:"member_id"`
}

// NewInitialAuthorizedInfo: AuthorizedInfoの新しいインスタンスを作成
func NewInitialAuthorizedInfo() AuthorizedInfo {
	return AuthorizedInfo{
		Role:     "anonymous",
		TenantID: "default",
		MemberID: "unknown",
	}
}

// LogEntry: ログエントリを保持する構造体
type LogEntry struct {
	Timestamp      time.Time       `json:"timestamp"`
	Level          string          `json:"level"`
	System         SystemInfo      `json:"system"`
	HTTP           HTTPRequestInfo `json:"http"`
	AuthorizedInfo AuthorizedInfo  `json:"authorized_info"`
	Error          string          `json:"error"`
	Message        string          `json:"message"`
}

// globalLogger: ロガーのグローバルインスタンス
var globalLogger *Logger

// New: ロガーの新しいインスタンスを作成
func New() *Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return &Logger{
		Logger: slog.New(handler),
	}
}

// GetLogger: ロガーのグローバルインスタンスを取得
func GetLogger() *Logger {
	if globalLogger == nil {
		globalLogger = New()
	}
	return globalLogger
}

// DebugContext: デバッグログを出力
func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.Logger.DebugContext(ctx, msg, args...)
}

// InfoContext: 情報ログを出力
func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.Logger.InfoContext(ctx, msg, args...)
}

// WarnContext: 警告ログを出力
func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.Logger.WarnContext(ctx, msg, args...)
}

// ErrorContext: エラーログを出力
func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.Logger.ErrorContext(ctx, msg, args...)
}

// FatalContext: 致命的なエラーログを出力
func (l *Logger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.Logger.ErrorContext(ctx, msg, args...)
	os.Exit(1)
}

// contextKey: コンテキストキー
type contextKey string

const (
	authorizedInfoKey contextKey = "authorized_info"
	requestIDKey      contextKey = "request_id"
	systemInfoKey     contextKey = "system_info"
)

// Context: コンテキストの型エイリアス
type Context = context.Context

// WithSystemInfo: コンテキストにシステム情報を設定
func WithSystemInfo(ctx context.Context, info SystemInfo) context.Context {
	return context.WithValue(ctx, systemInfoKey, info)
}

// GetSystemInfo: コンテキストからシステム情報を取得
func GetSystemInfo(ctx context.Context) (SystemInfo, bool) {
	info, ok := ctx.Value(systemInfoKey).(SystemInfo)
	return info, ok
}

// WithValue: コンテキストに値を設定
func WithValue(ctx context.Context, key interface{}, val interface{}) context.Context {
	return context.WithValue(ctx, key, val)
}

// WithAuthorizedInfo: コンテキストに認可情報を設定
func WithAuthorizedInfo(ctx context.Context, info AuthorizedInfo) context.Context {
	return context.WithValue(ctx, authorizedInfoKey, info)
}

// GetAuthorizedInfo: コンテキストから認可情報を取得
func GetAuthorizedInfo(ctx context.Context) (AuthorizedInfo, bool) {
	info, ok := ctx.Value(authorizedInfoKey).(AuthorizedInfo)
	return info, ok
}

// WithRequestID: コンテキストにリクエストIDを設定
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

// GenerateAndSetRequestID: UUIDを生成してコンテキストに設定
func GenerateAndSetRequestID(ctx context.Context) context.Context {
	requestID := uuid.New().String()
	return WithRequestID(ctx, requestID)
}

// GetRequestID: コンテキストからリクエストIDを取得
func GetRequestID(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(requestIDKey).(string)
	return requestID, ok
}

// HTTPRequest: HTTPリクエスト情報をログに出力
func (l *Logger) HTTPRequest(ctx context.Context, entry LogEntry) {
	args := []any{
		"timestamp", entry.Timestamp.Format(time.RFC3339),
		"level", entry.Level,
		"message", entry.Message,
		"system", entry.System,
		"http", entry.HTTP,
		"authorized_info", entry.AuthorizedInfo,
	}

	if entry.Error != "" {
		args = append(args, "error", entry.Error)
	}

	switch entry.Level {
	case "debug":
		l.DebugContext(ctx, entry.Message, args...)
	case "info":
		l.InfoContext(ctx, entry.Message, args...)
	case "warn":
		l.WarnContext(ctx, entry.Message, args...)
	case "error":
		l.ErrorContext(ctx, entry.Message, args...)
	case "fatal":
		l.FatalContext(ctx, entry.Message, args...)
	default:
		l.InfoContext(ctx, entry.Message, args...)
	}
}
