package routes

import (
	"net/http"
	"redditclone/internal/handler"
	"redditclone/internal/middleware"
)

func SetRoutes(h *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	authMux := http.NewServeMux()
	mux.HandleFunc("/", h.MainPage)
	mux.HandleFunc("/static/", h.Static)
	mux.HandleFunc("GET /api/posts/", h.GetPosts)
	mux.HandleFunc("POST /api/register", h.RegisterUser)
	mux.HandleFunc("POST /api/login", h.LoginUser)
	mux.HandleFunc("GET /api/posts/{category}", h.GetPostsByCategory)
	mux.HandleFunc("GET /api/post/{id}", h.GetPostById)
	mux.HandleFunc("GET /api/user/{login}", h.GetPostByUserLogin)
	authMux.HandleFunc("POST /api/post/{id}", h.CreateComment)
	authMux.HandleFunc("POST /api/posts", h.CreatePost)
	authMux.HandleFunc("GET /api/post/{post_id}/upvote", h.Upvote)
	authMux.HandleFunc("GET /api/post/{post_id}/downvote", h.Downvote)
	authMux.HandleFunc("GET /api/post/{post_id}/unvote", h.Unvote)
	authMux.HandleFunc("DELETE /api/post/{id}", h.DeletePostById)
	authMux.HandleFunc("DELETE /api/post/{post_id}/{comment_id}", h.DeleteCommentById)

	mux.Handle("/api/", middleware.Auth(authMux))

	return mux
}
