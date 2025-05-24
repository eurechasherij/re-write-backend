package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"re-write-backend/internal/auth"

	"github.com/golang-jwt/jwt/v5"
)

type key int

const UserCtxKey key = 0

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing auth token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			http.Error(w, "Invalid auth header", http.StatusUnauthorized)
			return
		}

		jwtKey := []byte(os.Getenv("JWT_SECRET"))
		token, err := jwt.ParseWithClaims(tokenStr, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			log.Printf("JWT parse error: %v", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*auth.Claims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserCtxKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
