package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Tarun-GH/go-rest-microservice/internal/repository"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandlers(w http.ResponseWriter, r *http.Request) {
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

	//DB lookup work
	user, err := repository.GetUserByEmail(DB, req.Email, "users")
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	//password check and authentication
	_ = user
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"login request received properly"}`))
}
