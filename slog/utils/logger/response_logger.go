package logger

import (
	"context"
	"net/http"
)

// ResponseWriterWrapper: wrapper to record status code and context
type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
	ctx        *context.Context
}

// NewResponseWriterWrapper: create new ResponseWriterWrapper
func NewResponseWriterWrapper(w http.ResponseWriter) *ResponseWriterWrapper {
	defaultStatusCode := http.StatusOK
	return &ResponseWriterWrapper{
		ResponseWriter: w,
		statusCode:     defaultStatusCode,
	}
}

// WriteHeader: write header to ResponseWriterWrapper
func (rw *ResponseWriterWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// Write: write data to ResponseWriterWrapper
func (rw *ResponseWriterWrapper) Write(data []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = 200
	}
	return rw.ResponseWriter.Write(data)
}

// UpdateContext: update context in ResponseWriterWrapper
func (rw *ResponseWriterWrapper) UpdateContext(ctx context.Context) {
	if rw.ctx == nil {
		rw.ctx = &ctx
	} else {
		*rw.ctx = ctx
	}
}

// GetContext: get context from ResponseWriterWrapper
func (rw *ResponseWriterWrapper) GetContext() *context.Context {
	return rw.ctx
}

// GetStatusCode: get status code from ResponseWriterWrapper
func (rw *ResponseWriterWrapper) GetStatusCode() int {
	return rw.statusCode
}
