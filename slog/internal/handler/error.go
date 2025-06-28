package handler

import (
	"net/http"
)

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// HandleError400 handles 400 Bad Request error
func HandleError400(w http.ResponseWriter, r *http.Request) {
	base := NewBaseHandler()
	response := ErrorResponse{
		Error:   "bad_request",
		Message: "Invalid request parameters",
	}
	base.WriteJSONResponse(w, r, http.StatusBadRequest, response)
}

// HandleError401 handles 401 Unauthorized error
func HandleError401(w http.ResponseWriter, r *http.Request) {
	base := NewBaseHandler()
	response := ErrorResponse{
		Error:   "unauthorized",
		Message: "Authentication required",
	}
	base.WriteJSONResponse(w, r, http.StatusUnauthorized, response)
}

// HandleError403 handles 403 Forbidden error
func HandleError403(w http.ResponseWriter, r *http.Request) {
	base := NewBaseHandler()
	response := ErrorResponse{
		Error:   "forbidden",
		Message: "Access denied",
	}
	base.WriteJSONResponse(w, r, http.StatusForbidden, response)
}

// HandleError404 handles 404 Not Found error
func HandleError404(w http.ResponseWriter, r *http.Request) {
	base := NewBaseHandler()
	response := ErrorResponse{
		Error:   "not_found",
		Message: "Resource not found",
	}
	base.WriteJSONResponse(w, r, http.StatusNotFound, response)
}

// HandleError500 handles 500 Internal Server Error
func HandleError500(w http.ResponseWriter, r *http.Request) {
	base := NewBaseHandler()
	response := ErrorResponse{
		Error:   "internal_server_error",
		Message: "Internal server error occurred",
	}
	base.WriteJSONResponse(w, r, http.StatusInternalServerError, response)
}
