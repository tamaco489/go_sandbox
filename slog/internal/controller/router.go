package controller

import (
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/internal/handler"
	"github.com/tamaco489/go_sandbox/slog/internal/middleware/auth"
	"github.com/tamaco489/go_sandbox/slog/internal/middleware/logging"
)

type Router struct {
	logRouter  *logging.LogRouter
	authorizer auth.Authorizer
}

func NewRouter() *Router {
	return &Router{
		logRouter:  logging.NewLogRouter(),
		authorizer: auth.NewAuth(),
	}
}

func (r *Router) RegisterRoutes() {
	r.logRouter.HandleFunc("/api/v1/health", handler.HandleHealth) // NOTE: Skip authorization for health check
	r.logRouter.HandleFunc("/api/v1/users/me", auth.WithAuth(r.authorizer, handler.HandleUserMe))
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.logRouter.ServeHTTP(w, req)
}
