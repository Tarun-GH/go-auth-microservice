package routes

import (
	"github.com/Tarun-GH/go-rest-microservice/internal/handlers"
	"github.com/Tarun-GH/go-rest-microservice/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *handlers.Handler) {
	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
	r.Post("/refresh", h.Refresh)
	r.Post("/logout", h.Logout)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(h.JWTSecret))
		r.Get("/protected", h.Protected)
	})
}
