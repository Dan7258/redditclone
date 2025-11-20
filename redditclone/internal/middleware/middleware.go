package middleware

import (
	"context"
	"net/http"
	"redditclone/internal/models"
	"redditclone/jwt"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.ParseToken(r)
		if err != nil || !token.Valid {
			http.Error(w, `{"message": "invalid token"}`, http.StatusUnauthorized)
			return
		}
		claims, err := jwt.ParseClaims(token)

		if err != nil {
			http.Error(w, `{"message": "invalid token claims"}`, http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserFromContext(ctx context.Context) (*models.User, bool) {
	user, ok := ctx.Value("user").(*models.User)
	return user, ok
}
