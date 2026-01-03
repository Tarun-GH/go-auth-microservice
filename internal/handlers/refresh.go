package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Tarun-GH/go-rest-microservice/internal/config"
	"github.com/Tarun-GH/go-rest-microservice/internal/utils"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	//verifying and loopup of refreshToken
	userID, ok := utils.GetUserIDFromRefreshToken(req.RefreshToken)
	if !ok {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	//deleting old refresh token (for rotation purposes)
	_ = utils.DeleteRefreshToken(req.RefreshToken)

	//generating new refreshToken as well
	newRefreshToken, err := utils.GenerateRefresh(userID)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	//generate new accessToken
	cfg := config.Load() //fetching config data
	accessToken, err := utils.GenerateToken([]byte(cfg.JWTSecret), userID, "")
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"new_access_token":  accessToken,
		"new_refresh_token": newRefreshToken,
	})
}
