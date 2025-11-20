package main

import (
	"log"
	"net/http"
	"redditclone/internal/config"
	"redditclone/internal/handler"
	"redditclone/internal/middleware"
	"redditclone/internal/repository"
	"redditclone/jwt"
	"time"
)

func main() {
	mux := http.NewServeMux()
	authMux := http.NewServeMux()
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
	fs := http.FileServer(http.Dir("./web/html"))
	mux.Handle("/", fs)
	staticHandler := http.FileServer(http.Dir("./web"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))
	mux.HandleFunc("GET /api/posts/", h.GetPosts)
	mux.HandleFunc("POST /api/register", h.RegisterUser)
	mux.HandleFunc("POST /api/login", h.LoginUser)
	mux.HandleFunc("GET /api/posts/{category}", h.GetPostsByCategory)
	mux.HandleFunc("GET /api/post/{id}", h.GetPostById)
	authMux.HandleFunc("POST /api/post/{id}", h.CreateComment)
	authMux.HandleFunc("POST /api/posts", h.CreatePost)
	authMux.HandleFunc("DELETE /api/post/{id}", h.DeletePostById)

	mux.Handle("/api/", middleware.Auth(authMux))
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
