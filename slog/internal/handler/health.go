package handler

import (
	"net/http"
)

// HealthResponse represents the health check response structure
type HealthResponse struct {
	Message string `json:"message"`
}

// HandleHealth handles health check API
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	base := NewBaseHandler()
	response := HealthResponse{Message: "ok"}
	base.WriteJSONResponse(w, r, http.StatusOK, response)
}
