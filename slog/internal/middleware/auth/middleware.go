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

		// 検証中のため意図的にエラーを返す
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return

		authInfo, err := authorizer.Authorize(r.Context(), r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// 認可情報をコンテキストに更新
		ctx := logger.WithAuthorizedInfo(r.Context(), *authInfo)
		r = r.WithContext(ctx)

		// 更新されたリクエストを次のハンドラーに渡す
		next.ServeHTTP(w, r)
	}
}
