package logger

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const (
	authorizedInfoKey contextKey = "authorized_info"
	requestIDKey      contextKey = "request_id"
	systemInfoKey     contextKey = "system_info"
	statusCodeKey     contextKey = "status_code"
)

// SetSystemInfoCtx: Set system information in context
func SetSystemInfoCtx(ctx context.Context, info SystemInfo) context.Context {
	return context.WithValue(ctx, systemInfoKey, info)
}

// GetSystemInfoCtx: Get system information from context
func GetSystemInfoCtx(ctx context.Context) (SystemInfo, bool) {
	info, ok := ctx.Value(systemInfoKey).(SystemInfo)
	return info, ok
}

// SetAuthorizedInfoCtx: Set authorized information in context
func SetAuthorizedInfoCtx(ctx context.Context, info AuthorizedInfo) context.Context {
	return context.WithValue(ctx, authorizedInfoKey, info)
}

// GetAuthorizedInfoCtx: Get authorized information from context
func GetAuthorizedInfoCtx(ctx context.Context) (AuthorizedInfo, bool) {
	info, ok := ctx.Value(authorizedInfoKey).(AuthorizedInfo)
	return info, ok
}

// SetRequestIDCtx: Set request ID in context
func SetRequestIDCtx(ctx context.Context) context.Context {
	requestID := uuid.New().String()
	return context.WithValue(ctx, requestIDKey, requestID)
}

// GetRequestIDCtx: Get request ID from context
func GetRequestIDCtx(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(requestIDKey).(string)
	return requestID, ok
}

// SetStatusCodeCtx: Set status code in context
func SetStatusCodeCtx(ctx context.Context, statusCode int) context.Context {
	return context.WithValue(ctx, statusCodeKey, statusCode)
}
