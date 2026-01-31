package middlewares

import (
	dbConfig "AuthInGo/config/db"
	env "AuthInGo/config/env"
	repo "AuthInGo/db/repositories"
	"context"
	"fmt"
	"net/http"
	"strconv"
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

		userIdFloat, okId := claims["id"].(float64)
		email, okEmail := claims["email"].(string)

		userId := int(userIdFloat)

		fmt.Println(userId, email)

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

// TODO: Will fix this later

// func RequireAllRoles(roles ...string) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 			userIdStr, ok := r.Context().Value("userID").(string)

// 			fmt.Println("RequireAllRoles : ", userIdStr)
// 			userId, err := strconv.ParseInt(userIdStr, 10, 64)
// 			if !ok || err != nil {
// 				http.Error(w, "User ID not found in context", http.StatusUnauthorized)
// 				return
// 			}

// 			dbConn, dbErr := dbConfig.SetupDB()

// 			if dbErr != nil {
// 				http.Error(w, "Database connection error"+dbErr.Error(), http.StatusInternalServerError)
// 				return
// 			}

// 			urr := repo.NewUserRolesRepository(dbConn)

// 			hasAllRoles, err := urr.HasAllRoles(userId, roles)
// 			if err != nil {
// 				http.Error(w, "Error checking user roles", http.StatusInternalServerError)
// 				return
// 			}

// 			if !hasAllRoles {
// 				http.Error(w, "User does not have all required roles", http.StatusForbidden)
// 				return
// 			}

// 			fmt.Println("User", userId, "has all required roles:", roles)
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

func RequireAllRoles(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userIDVal := r.Context().Value("userID")
			if userIDVal == nil {
				http.Error(w, "User ID not found in context", http.StatusUnauthorized)
				return
			}

			var userId int64

			switch v := userIDVal.(type) {
			case int:
				userId = int64(v)
			case int64:
				userId = v
			case float64: // JWT leak safety
				userId = int64(v)
			case string:
				parsed, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					http.Error(w, "Invalid user ID", http.StatusUnauthorized)
					return
				}
				userId = parsed
			default:
				http.Error(w, "Invalid user ID type", http.StatusUnauthorized)
				return
			}

			fmt.Println("RequireAllRoles userID:", userId)

			dbConn, dbErr := dbConfig.SetupDB()
			if dbErr != nil {
				http.Error(w, "Database connection error: "+dbErr.Error(), http.StatusInternalServerError)
				return
			}

			urr := repo.NewUserRolesRepository(dbConn)

			hasAllRoles, err := urr.HasAllRoles(userId, roles)
			if err != nil {
				http.Error(w, "Error checking user roles", http.StatusInternalServerError)
				return
			}

			if !hasAllRoles {
				http.Error(w, "User does not have all required roles", http.StatusForbidden)
				return
			}

			fmt.Println("User", userId, "has all required roles:", roles)
			next.ServeHTTP(w, r)
		})
	}
}

func RequireAnyRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userIDVal := r.Context().Value("userID")
			if userIDVal == nil {
				http.Error(w, "User ID not found in context", http.StatusUnauthorized)
				return
			}

			var userId int64

			switch v := userIDVal.(type) {
			case int:
				userId = int64(v)
			case int64:
				userId = v
			case float64: // JWT leak safety
				userId = int64(v)
			case string:
				parsed, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					http.Error(w, "Invalid user ID", http.StatusUnauthorized)
					return
				}
				userId = parsed
			default:
				http.Error(w, "Invalid user ID type", http.StatusUnauthorized)
				return
			}

			fmt.Println("RequireAnyRole userID:", userId)

			dbConn, dbErr := dbConfig.SetupDB()
			if dbErr != nil {
				http.Error(w, "Database connection error: "+dbErr.Error(), http.StatusInternalServerError)
				return
			}

			urr := repo.NewUserRolesRepository(dbConn)

			hasAnyRole, err := urr.HasAnyRole(userId, roles)
			if err != nil {
				http.Error(w, "Error checking user roles", http.StatusInternalServerError)
				return
			}

			if !hasAnyRole {
				http.Error(w, "User does not have any required role", http.StatusForbidden)
				return
			}

			fmt.Println("User", userId, "has any required role:", roles)
			next.ServeHTTP(w, r)
		})
	}
}
