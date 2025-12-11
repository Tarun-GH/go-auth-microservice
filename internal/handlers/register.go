package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Tarun-GH/go-rest-microservice/internal/repository"
	"github.com/Tarun-GH/go-rest-microservice/internal/utils"
	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	//Putting data into the RegisterReq. struct
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	//error handling
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		http.Error(w, "Name is required!", http.StatusBadRequest)
		return
	}
	if req.Email == "" || !strings.Contains(req.Email, "@") {
		http.Error(w, "Invalid Email!", http.StatusBadRequest)
		return
	}
	if len(req.Password) < 6 {
		http.Error(w, "Password cannot be less than 6 characters", http.StatusBadRequest)
		return
	}

	//Hashing the password
	hashed_pass, err := utils.HashPassword(req.Password)
	if err != nil {
		fmt.Println("Couldn't hash the password!", err)
		return
	}
	fmt.Println("Hashed Password:", hashed_pass)

	//repo se lekar call kiya insertion ka
	err = repository.InsertUser(DB, req.Name, req.Email, hashed_pass, "users")
	if err != nil {
		fmt.Println("Error inside the registerhander during inserting to DB:", err)
		return
	}

	//updating status
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"user registered"}`))

}
