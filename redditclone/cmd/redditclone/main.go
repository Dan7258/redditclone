package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	_ "net/http/pprof"
	"redditclone/internal/config"
	"redditclone/internal/handler"
	"redditclone/internal/middleware"
	"redditclone/internal/repository"
	"redditclone/internal/routes"
	"redditclone/pkg/jwt"
	"time"
)

func main() {
	middleware.Duration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "A histogram of latencies for requests.",
			Buckets: []float64{.25, .5, 1, 2.5, 5, 10},
		},
		[]string{"path", "method"},
	)
	prometheus.MustRegister(middleware.Duration)
	config.Init()
	err := jwt.Init()
	if err != nil {
		log.Fatal(err)
	}
	db := &repository.PostgresDB{}
	err = db.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	h := handler.InitHandler(db)
	mux := routes.SetRoutes(h)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      middleware.Logger(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	h.Logger.Info("Starting server on port 8080")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
