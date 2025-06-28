package handler

import (
	"net/http"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func HandleProductByID(w http.ResponseWriter, r *http.Request) {
	// NOTE: Intentionally returning 5xx
	http.Error(w, "server error", http.StatusInternalServerError)

	// base := NewBaseHandler()
	// productID := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
	// response := Product{
	// 	ID:   productID,
	// 	Name: "Product_ABC",
	// }
	// base.WriteJSONResponse(w, r, http.StatusOK, response)
}
