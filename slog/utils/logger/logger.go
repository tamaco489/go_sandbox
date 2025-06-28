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

// globalLogger: ロガーのグローバルインスタンス
var globalLogger *Logger

// New: ロガーの新しいインスタンスを作成
func New() *Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
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
	statusCodeKey     contextKey = "status_code"
)

// WithSystemInfo: コンテキストにシステム情報を設定
func WithSystemInfo(ctx context.Context, info SystemInfo) context.Context {
	return context.WithValue(ctx, systemInfoKey, info)
}

// GetSystemInfo: コンテキストからシステム情報を取得
func GetSystemInfo(ctx context.Context) (SystemInfo, bool) {
	info, ok := ctx.Value(systemInfoKey).(SystemInfo)
	return info, ok
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

// SetRequestID: UUIDを生成してコンテキストに設定
func SetRequestID(ctx context.Context) context.Context {
	requestID := uuid.New().String()
	return WithRequestID(ctx, requestID)
}

// GetRequestID: コンテキストからリクエストIDを取得
func GetRequestID(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(requestIDKey).(string)
	return requestID, ok
}

// WithStatusCode: コンテキストにHTTPステータスコードを設定
func WithStatusCode(ctx context.Context, statusCode int) context.Context {
	return context.WithValue(ctx, statusCodeKey, statusCode)
}

// ResponseWriterWrapper: ステータスコードとコンテキストを記録するラッパー
type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
	ctx        *context.Context
}

// NewResponseWriterWrapper: 新しいResponseWriterWrapperを作成
func NewResponseWriterWrapper(w http.ResponseWriter) *ResponseWriterWrapper {
	defaultStatusCode := http.StatusOK
	return &ResponseWriterWrapper{
		ResponseWriter: w,
		statusCode:     defaultStatusCode,
	}
}

func (rw *ResponseWriterWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *ResponseWriterWrapper) Write(data []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = 200
	}
	return rw.ResponseWriter.Write(data)
}

func (rw *ResponseWriterWrapper) UpdateContext(ctx context.Context) {
	if rw.ctx == nil {
		rw.ctx = &ctx
	} else {
		*rw.ctx = ctx
	}
}

func (rw *ResponseWriterWrapper) GetContext() *context.Context {
	return rw.ctx
}

func (rw *ResponseWriterWrapper) GetStatusCode() int {
	return rw.statusCode
}
