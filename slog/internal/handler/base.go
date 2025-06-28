package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

// BaseHandler: base handler
type BaseHandler struct{}

// NewBaseHandler: create new base handler
func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

// SetStatusCode: set status code to context
func (h *BaseHandler) SetStatusCode(r *http.Request, statusCode int) *http.Request {
	ctx := logger.SetStatusCodeContext(r.Context(), statusCode)
	return r.WithContext(ctx)
}

// WriteJSONResponse: write JSON response and set status code
func (h *BaseHandler) WriteJSONResponse(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	// set status code to context
	ctx := logger.SetStatusCodeContext(r.Context(), statusCode)
	_ = r.WithContext(ctx) // Update context but not used in this function

	jsonData, err := json.Marshal(data)
	if err != nil {
		// set status code to context for error case
		ctx = logger.SetStatusCodeContext(r.Context(), http.StatusInternalServerError)
		_ = r.WithContext(ctx) // Update context for error case
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// WriteErrorResponse: write error response and set status code
func (h *BaseHandler) WriteErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	// set status code to context
	r = h.SetStatusCode(r, statusCode)
	http.Error(w, message, statusCode)
}
