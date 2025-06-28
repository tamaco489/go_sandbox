package handler

import (
	"encoding/json"
	"net/http"
)

// HealthResponse represents the health check response structure
type HealthResponse struct {
	Message string `json:"message"`
}

// HandleHealth handles health check API
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	response := HealthResponse{Message: "ok"}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
