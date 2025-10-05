package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// JWTAuth is a middleware function that checks for a valid JWT in the Authorization header.
// It expects the token to be prefixed with "Bearer ". If the token is missing or invalid,
// it responds with an HTTP 401 Unauthorized status. If the token is valid, it calls the next handler.
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		envSecret := os.Getenv("JWT_SECRET")
		jwtSecret := []byte("gwi-jwt-secret")
		if envSecret != "" {
			jwtSecret = []byte(envSecret)
		}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
