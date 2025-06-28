package handler

import (
	"net/http"
)

// UserMeResponse represents the user me response structure
type UserMeResponse struct {
	UID string `json:"uid"`
}

// HandleUserMe handles user me API
func HandleUserMe(w http.ResponseWriter, r *http.Request) {
	base := NewBaseHandler()

	// For now, return a mock UID
	// In a real application, this would be extracted from authentication context
	response := UserMeResponse{
		UID: "864c857e-bc03-7b09-5b8f-750d312636c3",
	}

	base.WriteJSONResponse(w, r, http.StatusOK, response)
}
