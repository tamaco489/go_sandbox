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

func (rw *ResponseWriterWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *ResponseWriterWrapper) Write(data []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = 200
	}
	return rw.ResponseWriter.Write(data)
}

func (rw *ResponseWriterWrapper) UpdateContext(ctx context.Context) {
	if rw.ctx == nil {
		rw.ctx = &ctx
	} else {
		*rw.ctx = ctx
	}
}

func (rw *ResponseWriterWrapper) GetContext() *context.Context {
	return rw.ctx
}

func (rw *ResponseWriterWrapper) GetStatusCode() int {
	return rw.statusCode
}
