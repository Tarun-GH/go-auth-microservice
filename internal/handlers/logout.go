package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Tarun-GH/go-rest-microservice/internal/utils"
)

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	var req LogoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
	}

	//delete refresh_token
	_ = utils.DeleteRefreshToken(req.RefreshToken)

	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "logged out",
	})
}
