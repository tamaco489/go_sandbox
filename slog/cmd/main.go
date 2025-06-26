package main

import (
	"context"
	"log"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/internal/handler"
	"github.com/tamaco489/go_sandbox/slog/internal/logger"
)

func main() {
	port := ":8080"

	ctx := context.Background()
	logger.InfoContext(ctx, "Server started successfully", "port", port)

	router := handler.NewRouter()
	router.RegisterRoutes()

	if err := http.ListenAndServe(port, router); err != nil {
		logger.ErrorContext(ctx, "Server startup error", "error", err.Error())
		log.Fatal("Server startup error:", err)
	}
}
