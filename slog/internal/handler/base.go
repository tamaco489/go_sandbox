package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tamaco489/go_sandbox/slog/utils/logger"
)

// BaseHandler: ハンドラーのベースクラス
type BaseHandler struct{}

// NewBaseHandler: 新しいベースハンドラーを作成
func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

// SetStatusCode: ステータスコードをコンテキストに設定
func (h *BaseHandler) SetStatusCode(r *http.Request, statusCode int) *http.Request {
	ctx := logger.WithStatusCode(r.Context(), statusCode)
	return r.WithContext(ctx)
}

// WriteJSONResponse: JSONレスポンスを書き込み、ステータスコードを設定
func (h *BaseHandler) WriteJSONResponse(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	// ステータスコードをコンテキストに設定
	ctx := logger.WithStatusCode(r.Context(), statusCode)
	r = r.WithContext(ctx)

	jsonData, err := json.Marshal(data)
	if err != nil {
		// エラー時のステータスコードをコンテキストに設定
		ctx = logger.WithStatusCode(r.Context(), http.StatusInternalServerError)
		r = r.WithContext(ctx)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// WriteErrorResponse: エラーレスポンスを書き込み、ステータスコードを設定
func (h *BaseHandler) WriteErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	// エラー時のステータスコードをコンテキストに設定
	r = h.SetStatusCode(r, statusCode)
	http.Error(w, message, statusCode)
}
