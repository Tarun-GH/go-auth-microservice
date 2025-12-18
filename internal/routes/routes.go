package routes

import (
	"github.com/Tarun-GH/go-rest-microservice/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandlers)
}
