package auth

import (
	"context"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

type Authorizer interface {
	Authorize(ctx context.Context, r *http.Request) (*logger.AuthorizedInfo, error)
}

// ResponseWriterWrapper: コンテキストを更新できるラッパー
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

// WithAuth: 認可ミドルウェア
func WithAuth(authorizer Authorizer, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.GetLogger().DebugContext(r.Context(), "認可ミドルウェア開始", "path", r.URL.Path)

		// ResponseWriterWrapperかどうかチェック
		wrappedWriter, isWrapped := w.(*ResponseWriterWrapper)

		// 認可処理を実行
		authInfo, err := authorizer.Authorize(r.Context(), r)
		if err != nil {
			logger.GetLogger().DebugContext(r.Context(), "認可失敗", "error", err.Error())
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		logger.GetLogger().DebugContext(r.Context(), "認可成功", "auth_info", authInfo)

		// 認可情報をコンテキストに設定
		ctx := logger.WithAuthorizedInfo(r.Context(), *authInfo)
		r = r.WithContext(ctx)

		// ResponseWriterWrapperの場合はコンテキストも更新
		if isWrapped {
			wrappedWriter.UpdateContext(ctx)
		}

		logger.GetLogger().DebugContext(ctx, "認可情報をコンテキストに設定しました", "auth_info", authInfo)

		// 更新されたリクエストを次のハンドラーに渡す
		next.ServeHTTP(w, r)
	}
}
