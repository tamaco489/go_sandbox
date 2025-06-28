package handler

import (
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/internal/middleware"
)

type Router struct {
	mux        *http.ServeMux
	authorizer middleware.Authorizer
}

func NewRouter() *Router {
	return &Router{
		mux:        http.NewServeMux(),
		authorizer: middleware.NewAuth(),
	}
}

func (r *Router) RegisterRoutes() {
	r.mux.HandleFunc("/api/v1/health", middleware.LoggingMiddleware(HandleHealth))
	r.mux.HandleFunc("/api/v1/users/me", middleware.LoggingMiddleware(middleware.WithAuth(r.authorizer, HandleUserMe)))
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
