package logger

import (
	"context"

	"github.com/google/uuid"
)

// contextKey: コンテキストキー
type contextKey string

const (
	authorizedInfoKey contextKey = "authorized_info"
	requestIDKey      contextKey = "request_id"
	systemInfoKey     contextKey = "system_info"
	statusCodeKey     contextKey = "status_code"
)

// SetSystemInfoCtx: コンテキストにシステム情報を設定
func SetSystemInfoCtx(ctx context.Context, info SystemInfo) context.Context {
	return context.WithValue(ctx, systemInfoKey, info)
}

// GetSystemInfoCtx: コンテキストからシステム情報を取得
func GetSystemInfoCtx(ctx context.Context) (SystemInfo, bool) {
	info, ok := ctx.Value(systemInfoKey).(SystemInfo)
	return info, ok
}

// SetAuthorizedInfoCtx: コンテキストに認可情報を設定
func SetAuthorizedInfoCtx(ctx context.Context, info AuthorizedInfo) context.Context {
	return context.WithValue(ctx, authorizedInfoKey, info)
}

// GetAuthorizedInfoCtx: コンテキストから認可情報を取得
func GetAuthorizedInfoCtx(ctx context.Context) (AuthorizedInfo, bool) {
	info, ok := ctx.Value(authorizedInfoKey).(AuthorizedInfo)
	return info, ok
}

// SetRequestIDCtx: UUIDを生成してコンテキストに設定
func SetRequestIDCtx(ctx context.Context) context.Context {
	requestID := uuid.New().String()
	return context.WithValue(ctx, requestIDKey, requestID)
}

// GetRequestIDCtx: コンテキストからリクエストIDを取得
func GetRequestIDCtx(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(requestIDKey).(string)
	return requestID, ok
}

// SetStatusCodeCtx: コンテキストにHTTPステータスコードを設定
func SetStatusCodeCtx(ctx context.Context, statusCode int) context.Context {
	return context.WithValue(ctx, statusCodeKey, statusCode)
}
