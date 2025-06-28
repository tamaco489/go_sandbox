package main

import (
	"log"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/internal/controller"
	"github.com/tamaco489/go_sandbox/slog/internal/middleware/logging"
)

func main() {
	port := ":8080"
	router := controller.NewRouter()
	router.RegisterRoutes()

	// Apply global middleware
	handler := logging.RequestMiddleware(router.ServeHTTP)

	if err := http.ListenAndServe(port, http.HandlerFunc(handler)); err != nil {
		log.Fatal("Server startup error:", err)
	}
}
