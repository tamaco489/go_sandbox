package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tamaco489/go_sandbox/slog/internal/controller"
	"github.com/tamaco489/go_sandbox/slog/utils/configuration"
	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

// requestMiddleware: リクエストの開始と終了を管理するミドルウェア
func requestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// リクエスト開始時刻を記録
		startTime := time.Now()

		// リクエストIDを生成してコンテキストに設定
		ctx := logger.SetRequestIDCtx(r.Context())

		// システム情報を初期化
		env := configuration.GetEnvironment()
		systemInfo := logger.NewSystemInfo(env)
		ctx = logger.SetSystemInfoCtx(ctx, systemInfo)

		// 初期の認可情報を設定
		authInfo := logger.NewInitialAuthorizedInfo()
		ctx = logger.SetAuthorizedInfoCtx(ctx, authInfo)

		// 更新されたコンテキストでリクエストを更新
		r = r.WithContext(ctx)

		// ResponseWriterWrapperを作成（コンテキストのポインタを保持）
		wrappedWriter := logger.NewResponseWriterWrapper(w)

		// ctxフィールドを初期化してからUpdateContextを呼ぶ
		wrappedWriter.UpdateContext(ctx)

		// deferでリクエスト終了時のログ出力
		defer func() {
			// wrappedWriter.ctxから最新のコンテキストを取得（認可ミドルウェアで更新されたもの）
			finalAuthInfo, _ := logger.GetAuthorizedInfoCtx(*wrappedWriter.GetContext())
			systemInfo, _ := logger.GetSystemInfoCtx(*wrappedWriter.GetContext())
			requestID, _ := logger.GetRequestIDCtx(*wrappedWriter.GetContext())

			// 処理時間を計算
			latency := time.Since(startTime)

			// HTTP情報を作成
			httpInfo := logger.HTTPRequestInfo{
				Method:     r.Method,
				Path:       r.URL.Path,
				Status:     wrappedWriter.GetStatusCode(),
				Latency:    latency.String(),
				UserAgent:  r.UserAgent(),
				Referer:    r.Referer(),
				RemoteAddr: r.RemoteAddr,
				RequestID:  requestID,
			}

			// ステータスコードに応じてログレベルを決定
			statusCode := wrappedWriter.GetStatusCode()
			switch {
			// status: 5xx, level: error
			case statusCode >= http.StatusInternalServerError:
				logger.GetLogger().ErrorContext(*wrappedWriter.GetContext(), "Request completed",
					"status_code", statusCode,
					"http_info", httpInfo,
					"system_info", systemInfo,
					"auth_info", finalAuthInfo,
				)
			// status: 4xx, level: warn
			case statusCode >= http.StatusBadRequest:
				logger.GetLogger().WarnContext(*wrappedWriter.GetContext(), "Request completed",
					"status_code", statusCode,
					"http_info", httpInfo,
					"system_info", systemInfo,
					"auth_info", finalAuthInfo,
				)
			// status: 2xx, level: info
			case statusCode >= http.StatusOK && statusCode < http.StatusBadRequest:
				logger.GetLogger().InfoContext(*wrappedWriter.GetContext(), "Request completed",
					"status_code", statusCode,
					"http_info", httpInfo,
					"system_info", systemInfo,
					"auth_info", finalAuthInfo,
				)
			default:
				logger.GetLogger().InfoContext(*wrappedWriter.GetContext(), "Request completed",
					"status_code", statusCode,
					"http_info", httpInfo,
					"system_info", systemInfo,
					"auth_info", finalAuthInfo,
				)
			}
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
