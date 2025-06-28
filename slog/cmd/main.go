package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/tamaco489/go_sandbox/slog/internal/controller"
	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

// ResponseWriterWrapper: ステータスコードを記録するラッパー
type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
	ctx        *context.Context
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
	*rw.ctx = ctx
}

// requestMiddleware: リクエストの開始と終了を管理するミドルウェア
func requestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// リクエスト開始時刻を記録
		startTime := time.Now()

		// リクエストIDを生成してコンテキストに設定
		ctx := logger.SetRequestID(r.Context())

		// システム情報を初期化
		env := "dev" // TODO: 設定から取得
		systemInfo := logger.NewSystemInfo(env)

		// 初期の認可情報を設定
		authInfo := logger.NewInitialAuthorizedInfo()
		ctx = logger.WithAuthorizedInfo(ctx, authInfo)
		ctx = logger.WithSystemInfo(ctx, systemInfo)

		// 更新されたコンテキストでリクエストを更新
		r = r.WithContext(ctx)

		// ResponseWriterWrapperを作成（コンテキストのポインタを保持）
		wrappedWriter := &ResponseWriterWrapper{
			ResponseWriter: w,
			ctx:            &ctx,
		}

		// deferでリクエスト終了時のログ出力
		defer func() {
			// 最新のコンテキストから認可情報を取得
			finalAuthInfo, _ := logger.GetAuthorizedInfo(*wrappedWriter.ctx)
			systemInfo, _ := logger.GetSystemInfo(*wrappedWriter.ctx)
			requestID, _ := logger.GetRequestID(*wrappedWriter.ctx)

			// 処理時間を計算
			latency := time.Since(startTime)

			// HTTP情報を作成
			httpInfo := logger.HTTPRequestInfo{
				Method:     r.Method,
				Path:       r.URL.Path,
				Status:     wrappedWriter.statusCode,
				Latency:    latency.String(),
				UserAgent:  r.UserAgent(),
				Referer:    r.Referer(),
				RemoteAddr: r.RemoteAddr,
				RequestID:  requestID,
			}

			// ログ出力
			logger.GetLogger().InfoContext(*wrappedWriter.ctx, "Request completed",
				"status_code", wrappedWriter.statusCode,
				"http_info", httpInfo,
				"system_info", systemInfo,
				"auth_info", finalAuthInfo,
			)
		}()

		// 次のハンドラーを実行
		next.ServeHTTP(wrappedWriter, r)
	}
}

func main() {
	port := ":8080"
	router := controller.NewRouter()
	router.RegisterRoutes()

	// グローバルミドルウェアを適用
	handler := requestMiddleware(router.ServeHTTP)

	if err := http.ListenAndServe(port, http.HandlerFunc(handler)); err != nil {
		log.Fatal("Server startup error:", err)
	}
}
