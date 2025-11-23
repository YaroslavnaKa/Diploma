package api

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pass := os.Getenv("TODO_PASSWORD")

		if len(pass) == 0 {
			next(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			// Куки нет -> 401
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		jwtToken := cookie.Value
		token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid claims", http.StatusUnauthorized)
			return
		}

		if claims["password_hash"] != hash(pass) {
			http.Error(w, "Token outdated", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
