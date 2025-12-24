package routes

import (
	"github.com/Tarun-GH/go-rest-microservice/internal/handlers"
	"github.com/Tarun-GH/go-rest-microservice/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandlers)
	r.Post("/refresh", handlers.RefreshHandler)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/protected", handlers.ProtectedHandler)
	})

}
