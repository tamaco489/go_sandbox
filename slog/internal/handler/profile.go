package handler

import "net/http"

// UserProfileMeResponse represents the user profile me response structure
type UserProfileMeResponse struct {
	UID       string `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// HandleUserProfileMe handles user profile me API
func HandleUserProfileMe(w http.ResponseWriter, r *http.Request) {
	// NOTE: Intentionally returning 4xx
	http.Error(w, "Incorrect request", http.StatusBadRequest)

	// base := NewBaseHandler()

	// response := UserProfileMeResponse{
	// 	UID:       "864c857e-bc03-7b09-5b8f-750d312636c3",
	// 	FirstName: "John",
	// 	LastName:  "Doe",
	// 	Email:     "john.doe@example.com",
	// 	Phone:     "1234567890",
	// }
	// base.WriteJSONResponse(w, r, http.StatusOK, response)
}
