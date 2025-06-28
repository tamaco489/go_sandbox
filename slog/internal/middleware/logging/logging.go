package logging

import (
	"net/http"
	"time"

	"github.com/tamaco489/go_sandbox/slog/utils/configuration"
	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

// LogRouter: ログ処理を含むカスタムルーター
type LogRouter struct {
	mux *http.ServeMux
}

// NewLogRouter: 新しいログルーターを作成
func NewLogRouter() *LogRouter {
	return &LogRouter{
		mux: http.NewServeMux(),
	}
}

// HandleFunc: ハンドラーを登録
func (lr *LogRouter) HandleFunc(pattern string, handler http.HandlerFunc) {
	lr.mux.HandleFunc(pattern, handler)
}

// ServeHTTP: HTTPリクエストを処理し、ログを出力
func (lr *LogRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 元のServeHTTPを実行（ログ出力はWithLoggingで行うため、ここではログ出力しない）
	lr.mux.ServeHTTP(w, r)
}

// WithLogging: ログミドルウェア関数
func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// リクエスト開始時刻を記録
		startTime := time.Now()

		// 既存のコンテキストを使う
		ctx := r.Context()

		// システム情報を初期化
		env := configuration.GetEnvironment()
		systemInfo := logger.NewSystemInfo(env)

		// すでに認可情報がセットされていればそれを使う
		authInfo, ok := logger.GetAuthorizedInfo(ctx)
		if !ok {
			authInfo = logger.NewInitialAuthorizedInfo()
			ctx = logger.WithAuthorizedInfo(ctx, authInfo)
		}

		ctx = logger.WithSystemInfo(ctx, systemInfo)
		ctx = logger.SetRequestID(ctx) // 必要ならここでリクエストIDをセット

		// 更新されたコンテキストでリクエストを更新
		r = r.WithContext(ctx)

		// レスポンスライターをラップしてステータスコードをキャプチャ
		wrappedWriter := logger.NewResponseWriterWrapper(w, ctx, r, startTime, systemInfo)

		// 次のハンドラーを実行
		next.ServeHTTP(wrappedWriter, r)

		// ハンドラー実行後にログ出力（認可情報を含む）
		statusCode := wrappedWriter.GetStatusCode()
		
		// 現在のリクエストのコンテキストから最新の認可情報を取得
		// 注意: この時点でr.Context()は認可ミドルウェアで更新されたコンテキストを含んでいる
		currentReq := r
		finalAuthInfo, ok := logger.GetAuthorizedInfo(currentReq.Context())
		if !ok {
			finalAuthInfo = authInfo
		}

		// HTTP情報を作成
		httpInfo := logger.NewInitialHTTPRequestInfo(currentReq, startTime, statusCode)

		// ログレベルに応じて出力
		switch {
		// status: 5xx, level: error
		case statusCode >= http.StatusInternalServerError:
			logger.GetLogger().ErrorContext(currentReq.Context(), "Request completed",
				"status_code", statusCode,
				"http_info", httpInfo,
				"system_info", systemInfo,
				"auth_info", finalAuthInfo,
			)

		// status: 4xx, level: warn
		case statusCode >= http.StatusBadRequest:
			logger.GetLogger().WarnContext(currentReq.Context(), "Request completed",
				"status_code", statusCode,
				"http_info", httpInfo,
				"system_info", systemInfo,
				"auth_info", finalAuthInfo,
			)

		// status: 2xx, level: info
		case statusCode >= http.StatusOK && statusCode < http.StatusBadRequest:
			logger.GetLogger().InfoContext(currentReq.Context(), "Request completed",
				"status_code", statusCode,
				"http_info", httpInfo,
				"system_info", systemInfo,
				"auth_info", finalAuthInfo,
			)

		default:
			logger.GetLogger().InfoContext(currentReq.Context(), "Request completed",
				"status_code", statusCode,
				"http_info", httpInfo,
				"system_info", systemInfo,
				"auth_info", finalAuthInfo,
			)
		}
	}
}
