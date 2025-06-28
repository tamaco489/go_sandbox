package middleware

import (
	"net/http"
	"time"

	"github.com/tamaco489/go_sandbox/slog/utils/configuration"
	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

// LoggingMiddleware: HTTPリクエストのログ出力を行うミドルウェア
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// レスポンスをキャプチャするためのラッパー
		responseWriter := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// UUIDを生成してコンテキストに設定
		ctx := logger.GenerateAndSetRequestID(r.Context())

		// システム情報の設定
		systemInfo := logger.NewSystemInfo(configuration.GetEnvironment())

		// 認可情報の設定
		authInfo := logger.NewInitialAuthorizedInfo()

		// コンテキストにログ情報を設定
		ctx = logger.WithSystemInfo(ctx, systemInfo)
		ctx = logger.WithAuthorizedInfo(ctx, authInfo)
		r = r.WithContext(ctx)

		// 次のハンドラーを実行
		next.ServeHTTP(responseWriter, r)

		// HTTPリクエスト情報の設定
		httpInfo := logger.NewInitialHTTPRequestInfo(r, start, responseWriter.statusCode)

		// コンテキストにHTTP情報を追加
		ctx = logger.WithHTTPInfo(ctx, httpInfo)

		// ログレベルの決定
		level := "info"
		if responseWriter.statusCode >= 400 {
			level = "error"
		}

		// コンテキスト情報付きでログ出力
		logger.LogWithContext(ctx, level, "HTTP Request", logger.NewLogArgs())
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
