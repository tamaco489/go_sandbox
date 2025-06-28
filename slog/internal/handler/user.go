package handler

import (
	"encoding/json"
	"net/http"
)

// UserMeResponse represents the user me response structure
type UserMeResponse struct {
	UID string `json:"uid"`
}

// HandleUserMe handles user me API
func HandleUserMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// For now, return a mock UID
	// In a real application, this would be extracted from authentication context
	response := UserMeResponse{UID: "864c857e-bc03-7b09-5b8f-750d312636c3"}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
