package middleware

import (
	"context"
	"net/http"
)

type Authorizer interface {
	Authorize(ctx context.Context, r *http.Request) error
}

func WithAuth(authorizer Authorizer, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authorizer.Authorize(r.Context(), r); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
