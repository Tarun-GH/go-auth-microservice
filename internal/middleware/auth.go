package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Tarun-GH/go-rest-microservice/internal/utils"
)

type contextKey string

const userContextKey contextKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//---AuthMiddlerware verifies JWT before allowing access to protected routes
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" { //---missing token
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" { //---invalid token format
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := utils.VerifyToken(parts[1])
		if err != nil { //---invalid token
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
