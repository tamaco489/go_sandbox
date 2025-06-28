package logging

import (
	"log/slog"
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
	// リクエスト開始時刻を記録
	startTime := time.Now()

	// リクエストIDを生成してコンテキストに設定
	ctx := logger.SetRequestID(r.Context())

	// システム情報を初期化
	env := configuration.GetEnvironment()
	systemInfo := logger.NewSystemInfo(env)

	// 認可前のデフォルト認可情報
	authInfo := logger.NewInitialAuthorizedInfo()

	// コンテキストにリクエスト情報を設定
	ctx = logger.WithSystemInfo(ctx, systemInfo)
	ctx = logger.WithAuthorizedInfo(ctx, authInfo)

	// 更新されたコンテキストでリクエストを更新
	r = r.WithContext(ctx)

	// レスポンスライターをラップしてステータスコードをキャプチャ
	wrappedWriter := logger.NewResponseWriterWrapper(w, ctx)

	// 元のServeHTTPを実行
	lr.mux.ServeHTTP(wrappedWriter, r)

	// キャプチャされたステータスコードをコンテキストに設定
	statusCode := wrappedWriter.GetStatusCode()
	ctx = logger.WithStatusCode(ctx, statusCode)

	// 更新されたコンテキストでリクエストを更新
	r = r.WithContext(ctx)

	// ハンドラー実行後にログ出力
	logHTTPRequest(r, startTime, systemInfo, authInfo, statusCode)
}

// logHTTPRequest: HTTPリクエストのログを出力
func logHTTPRequest(r *http.Request, startTime time.Time, systemInfo logger.SystemInfo, authInfo logger.AuthorizedInfo, statusCode int) {
	// コンテキストから最新の認可情報を取得
	if latestAuthInfo, ok := logger.GetAuthorizedInfo(r.Context()); ok {
		authInfo = latestAuthInfo
	}

	// コンテキストからステータスコードを取得（フォールバックとして引数を使用）
	if ctxStatusCode, ok := logger.GetStatusCode(r.Context()); ok {
		statusCode = ctxStatusCode
	}

	// HTTP情報を作成
	httpInfo := logger.NewInitialHTTPRequestInfo(r, startTime, statusCode)

	switch {
	// status: 5xx, level: error
	case statusCode >= http.StatusInternalServerError:
		slog.ErrorContext(r.Context(), "Request completed",
			"status_code", statusCode,
			"http_info", httpInfo,
			"system_info", systemInfo,
			"auth_info", authInfo,
		)

	// status: 4xx, level: warn
	case statusCode >= http.StatusBadRequest:
		slog.WarnContext(r.Context(), "Request completed",
			"status_code", statusCode,
			"http_info", httpInfo,
			"system_info", systemInfo,
			"auth_info", authInfo,
		)

	// status: 2xx, level: info
	case statusCode >= http.StatusOK && statusCode < http.StatusBadRequest:
		slog.InfoContext(r.Context(), "Request completed",
			"status_code", statusCode,
			"http_info", httpInfo,
			"system_info", systemInfo,
			"auth_info", authInfo,
		)

	default:
		slog.InfoContext(r.Context(), "Request completed",
			"status_code", statusCode,
			"http_info", httpInfo,
			"system_info", systemInfo,
			"auth_info", authInfo,
		)
	}
}
