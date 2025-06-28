package main

import (
	"log"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/internal/controller"
)

func main() {
	port := ":8080"
	router := controller.NewRouter()
	router.RegisterRoutes()
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Server startup error:", err)
	}
}
