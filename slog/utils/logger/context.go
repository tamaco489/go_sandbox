package logger

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const (
	requestIDKey      contextKey = "request_id"
	systemInfoKey     contextKey = "system_info"
	authorizedInfoKey contextKey = "authorized_info"
	statusCodeKey     contextKey = "status_code"
)

// SetSystemInfoContext: Set system information in context
func SetSystemInfoContext(ctx context.Context, info SystemInfo) context.Context {
	return context.WithValue(ctx, systemInfoKey, info)
}

// GetSystemInfoContext: Get system information from context
func GetSystemInfoContext(ctx context.Context) (SystemInfo, bool) {
	info, ok := ctx.Value(systemInfoKey).(SystemInfo)
	return info, ok
}

// SetAuthorizedInfoContext: Set authorized information in context
func SetAuthorizedInfoContext(ctx context.Context, info AuthorizedInfo) context.Context {
	return context.WithValue(ctx, authorizedInfoKey, info)
}

// GetAuthorizedInfoContext: Get authorized information from context
func GetAuthorizedInfoContext(ctx context.Context) (AuthorizedInfo, bool) {
	info, ok := ctx.Value(authorizedInfoKey).(AuthorizedInfo)
	return info, ok
}

// SetRequestIDContext: Set request ID in context
func SetRequestIDContext(ctx context.Context) context.Context {
	requestID := uuid.New().String()
	return context.WithValue(ctx, requestIDKey, requestID)
}

// GetRequestIDContext: Get request ID from context
func GetRequestIDContext(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(requestIDKey).(string)
	return requestID, ok
}

// SetStatusCodeContext: Set status code in context
func SetStatusCodeContext(ctx context.Context, statusCode int) context.Context {
	return context.WithValue(ctx, statusCodeKey, statusCode)
}
