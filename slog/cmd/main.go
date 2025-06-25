package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/internal/handler"
)

func main() {
	// Server configuration
	port := ":8080"

	// Routing setup
	http.HandleFunc("/api/v1/health", handler.HandleHealth)

	// Start server
	slog.Info("Server started", "port", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server startup error:", err)
	}
}
