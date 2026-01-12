package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) Protected(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := map[string]string{"message": "authenticated is complete"}
	json.NewEncoder(w).Encode(resp)
}
