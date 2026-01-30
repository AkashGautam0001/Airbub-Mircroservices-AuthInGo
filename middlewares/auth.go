package middlewares

import (
	env "AuthInGo/config/env"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer") {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GetString("JWT_SECRET", "secret")), nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		userId, okId := claims["id"].(string)
		email, okEmail := claims["email"].(string)

		if !okId || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		fmt.Println("Authenticated user ID:", userId, "Email:", email)

		ctx := context.WithValue(r.Context(), "userID", userId)
		ctx = context.WithValue(ctx, "email", email)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
