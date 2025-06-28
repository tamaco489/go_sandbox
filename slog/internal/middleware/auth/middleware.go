package auth

import (
	"context"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

type Authorizer interface {
	Authorize(ctx context.Context, r *http.Request) (*logger.AuthorizedInfo, error)
}

// WithAuth: 認可ミドルウェア
func WithAuth(authorizer Authorizer, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.GetLogger().DebugContext(r.Context(), "認可ミドルウェア開始", "path", r.URL.Path)

		// ResponseWriterWrapperかどうかチェック
		wrappedWriter, isWrapped := w.(*logger.ResponseWriterWrapper)

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

		// リクエストのコンテキストを更新
		r = r.WithContext(ctx)

		// ResponseWriterWrapperの場合はコンテキストも更新
		if isWrapped {
			logger.GetLogger().DebugContext(ctx, "ResponseWriterWrapperのコンテキストを更新します", "auth_info", authInfo)
			wrappedWriter.UpdateContext(ctx)
		} else {
			logger.GetLogger().DebugContext(ctx, "ResponseWriterWrapperではありません", "auth_info", authInfo)
		}

		logger.GetLogger().DebugContext(ctx, "認可情報をコンテキストに設定しました", "auth_info", authInfo)

		// 更新されたリクエストを次のハンドラーに渡す
		next.ServeHTTP(w, r)
	}
}
