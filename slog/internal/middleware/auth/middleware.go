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
		// 認可処理を実行
		authInfo, err := authorizer.Authorize(r.Context(), r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// 認可情報をコンテキストに更新
		ctx := logger.WithAuthorizedInfo(r.Context(), *authInfo)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}
