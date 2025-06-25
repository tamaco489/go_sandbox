package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/internal/handler"
)

func main() {
	port := ":8080"

	router := handler.NewRouter()
	router.RegisterRoutes()

	slog.Info("Server started", "port", port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Server startup error:", err)
	}
}
