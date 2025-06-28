package middleware

import (
	"context"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

type Authorizer interface {
	Authorize(ctx context.Context, r *http.Request) (*logger.AuthorizedInfo, error)
}

func WithAuth(authorizer Authorizer, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authInfo, err := authorizer.Authorize(r.Context(), r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// 認可情報をコンテキストに設定
		ctx := logger.WithAuthorizedInfo(r.Context(), *authInfo)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}
