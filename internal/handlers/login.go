package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Tarun-GH/go-rest-microservice/internal/repository"
	"github.com/Tarun-GH/go-rest-microservice/internal/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	//creating a LoginRequest and putting data
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request:", http.StatusBadRequest)
		return
	}

	//simple checkup
	if req.Email == "" || req.Password == "" {
		http.Error(w, "Invalid email and password", http.StatusBadRequest)
		return
	}

	//DB lookup work & fetching
	dbUser, err := repository.GetUserByEmail(h.DB, req.Email, "users")
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	//password check and authentication
	if ok := utils.CheckPassword(req.Password, dbUser.PasswordHash); !ok {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	//Generate access token
	accessToken, err := utils.GenerateToken([]byte(h.JWTSecret), dbUser.ID, dbUser.Email)
	if err != nil {
		http.Error(w, "Couldn't generate token", http.StatusInternalServerError)
		return
	}

	//generate refresh token
	refreshToken, err := utils.GenerateRefresh(dbUser.ID)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusInternalServerError)
		return
	}

	//---Response
	w.Header().Set("Content-Type", "application/json") //---this Stays above both manual and auto parsing of string to []byte of json ^
	json.NewEncoder(w).Encode(map[string]string{       //Encode to json then sends it to the destination 'w' {http.ResponseWriter}
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
