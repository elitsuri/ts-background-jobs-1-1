package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/example/ts-background-jobs-1/pkg/jwt"
	"github.com/example/ts-background-jobs-1/pkg/response"
)

type contextKey string
const UserIDKey contextKey = "userID"

// Auth validates the JWT Bearer token and injects the user ID into context.
func Auth(token *jwt.Manager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			response.Error(w, http.StatusUnauthorized, "missing or invalid authorization header")
			return
		}
		claims, err := token.Verify(strings.TrimPrefix(auth, "Bearer "))
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "invalid or expired token")
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserID extracts the user ID from context (set by Auth middleware).
func GetUserID(r *http.Request) (int64, bool) {
	id, ok := r.Context().Value(UserIDKey).(int64)
	return id, ok
}
