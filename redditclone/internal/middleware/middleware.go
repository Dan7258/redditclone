package middleware

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
	"net/http"
	"os"
	"redditclone/pkg/jwt"
	"time"
)

var Duration *prometheus.HistogramVec

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

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Info(r.RequestURI)
	})
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		Duration.With(prometheus.Labels{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Observe(time.Since(start).Seconds())
	})
}
