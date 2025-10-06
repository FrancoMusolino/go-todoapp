package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/FrancoMusolino/go-todoapp/internal/domain/interfaces"
	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(userRepo interfaces.IUserRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.WriteError(w, http.StatusUnauthorized, "Missing Authorization header", nil)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) < 2 || strings.ToLower(parts[0]) != "bearer" {
				utils.WriteError(w, http.StatusUnauthorized, "Invalid Authorization header", nil)
				return
			}

			tokenStr := parts[1]
			if tokenStr == "" {
				utils.WriteError(w, http.StatusUnauthorized, "Missing Token on header", nil)
				return
			}

			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}

				return []byte(utils.GetEnvOrDefault("JWT_SECRET", "")), nil
			})

			if err != nil || !token.Valid {
				utils.WriteError(w, http.StatusUnauthorized, "invalid auth token", nil)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				utils.WriteError(w, http.StatusUnauthorized, "invalid token claims", nil)
				return
			}

			fmt.Println(claims)
			userID, ok := claims["id"].(string)
			if !ok {
				utils.WriteError(w, http.StatusUnauthorized, "invalid user ID in token", nil)
				return
			}

			username, ok := claims["username"].(string)
			if !ok {
				utils.WriteError(w, http.StatusUnauthorized, "invalid username in token", nil)
				return
			}

			user, _ := userRepo.GetByUsername(username)
			if !user.IsVerified() {
				utils.WriteError(w, http.StatusUnauthorized, "user not verified", nil)
				return

			}

			ctx := context.WithValue(r.Context(), "userID", userID)
			ctx = context.WithValue(r.Context(), "username", username)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
