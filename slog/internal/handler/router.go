package handler

import (
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) RegisterRoutes() {
	r.mux.HandleFunc("/api/v1/health", HandleHealth)
	r.mux.HandleFunc("/api/v1/users/me", HandleUserMe)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
