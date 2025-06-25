package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	// Server configuration
	port := ":8080"

	// Routing setup
	http.HandleFunc("/api/v1/health", handleHealth)

	// Start server
	slog.Info("Server started", "port", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server startup error:", err)
	}
}

// HealthResponse represents the health check response structure
type HealthResponse struct {
	Message string `json:"message"`
}

// handleHealth handles health check API
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	response := HealthResponse{Message: "ok"}
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
