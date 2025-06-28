package controller

import (
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/internal/handler"
	"github.com/tamaco489/go_sandbox/slog/internal/middleware/auth"
)

type Router struct {
	mux        *http.ServeMux
	authorizer auth.Authorizer
}

func NewRouter() *Router {
	return &Router{
		mux:        http.NewServeMux(),
		authorizer: auth.NewAuth(),
	}
}

func (r *Router) RegisterRoutes() {
	// ヘルスチェックは認可不要
	r.mux.HandleFunc("/api/v1/health", handler.HandleHealth)

	// ユーザー情報は認可が必要
	r.mux.HandleFunc("/api/v1/users/me", auth.WithAuth(r.authorizer, handler.HandleUserMe))
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
