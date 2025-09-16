package auth

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const ContextKey ctxKey = "jwtClaims"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if !strings.HasPrefix(h, "Bearer") {
			next.ServeHTTP(w, r)
			return
		}
		tokenStr := strings.TrimPrefix(h, "Bearer")
		token, _ := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if token != nil && token.Valid {
			if claims, ok := token.Claims.(*Claims); ok {
				r = r.WithContext(context.WithValue(r.Context(), ContextKey, claims))
			}
		}
		next.ServeHTTP(w, r)
	})
}
