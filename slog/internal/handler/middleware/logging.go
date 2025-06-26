package middleware

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/tamaco489/go_sandbox/slog/internal/logger"
)

// LoggingMiddleware: HTTPリクエストのログ出力を行うミドルウェア
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// レスポンスをキャプチャするためのラッパー
		responseWriter := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// リクエストIDの生成（簡易版）
		requestID := generateRequestID()

		// コンテキストにリクエストIDを追加
		ctx := context.WithValue(r.Context(), "request_id", requestID)
		r = r.WithContext(ctx)

		// 次のハンドラーを実行
		next.ServeHTTP(responseWriter, r)

		// 処理時間の計算
		duration := time.Since(start)

		// システム情報の設定
		hostname, _ := os.Hostname()
		systemInfo := logger.SystemInfo{
			Environment: getEnvironment(),
			Service:     "slog-server",
			Hostname:    hostname,
			AccountID:   "system",
		}

		// HTTPリクエスト情報の設定
		httpInfo := logger.HTTPRequestInfo{
			Method:     r.Method,
			Path:       r.URL.Path,
			Status:     responseWriter.statusCode,
			Latency:    duration.String(),
			UserAgent:  r.UserAgent(),
			Referer:    r.Referer(),
			RemoteAddr: r.RemoteAddr,
			RequestID:  requestID,
		}

		// 認可情報の設定（簡易版）
		authInfo := logger.AuthorizedInfo{
			Role:     "anonymous",
			TenantID: "default",
			MemberID: "unknown",
		}

		// ログレベルの決定
		level := "info"
		if responseWriter.statusCode >= 400 {
			level = "error"
		}

		// ログエントリの作成
		entry := logger.LogEntry{
			Timestamp:      time.Now(),
			Level:          level,
			System:         systemInfo,
			HTTP:           httpInfo,
			AuthorizedInfo: authInfo,
			Message:        "HTTP Request",
		}

		// ログ出力
		logger.HTTPRequest(ctx, entry)
	}
}

// responseWriter: レスポンスのステータスコードをキャプチャするためのラッパー
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// generateRequestID: リクエストIDを生成（簡易版）
func generateRequestID() string {
	return time.Now().Format("20060102150405") + "-" + time.Now().Format("000000000")
}

// getEnvironment: 環境変数から環境を取得
func getEnvironment() string {
	if env := os.Getenv("ENV"); env != "" {
		return env
	}
	return "dev"
}
