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
	// health check
	r.mux.HandleFunc("/api/v1/health", handler.HandleHealth)

	// public routes
	r.mux.HandleFunc("/api/v1/products/{id}", handler.HandleProductByID)

	// pricate routes
	r.mux.HandleFunc("/api/v1/users/me", auth.WithAuth(r.authorizer, handler.HandleUserMe))
	r.mux.HandleFunc("/api/v1/users/profile/me", auth.WithAuth(r.authorizer, handler.HandleUserProfileMe))
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
