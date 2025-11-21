package main

import (
	"log"
	"net/http"
	"redditclone/internal/config"
	"redditclone/internal/handler"
	"redditclone/internal/repository"
	"redditclone/internal/routes"
	"redditclone/jwt"
	"time"
)

func main() {
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
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
