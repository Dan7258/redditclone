package routes

import (
	"net/http"
	"redditclone/internal/handler"
	"redditclone/internal/middleware"
)

func SetRoutes(h *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	authMux := http.NewServeMux()
	mux.HandleFunc("/", h.MainPage)                                                    // главная страница
	mux.HandleFunc("/static/", h.Static)                                               // подключение статики
	mux.HandleFunc("GET /api/posts/", h.GetPosts)                                      // получение всех постов
	mux.HandleFunc("POST /api/register", h.RegisterUser)                               // регистрация пользователя
	mux.HandleFunc("POST /api/login", h.LoginUser)                                     // авторизация пользователя
	mux.HandleFunc("GET /api/posts/{category}", h.GetPostsByCategory)                  // получение списка пользователей по категории
	mux.HandleFunc("GET /api/post/{id}", h.GetPostByID)                                // получение поста по id
	mux.HandleFunc("GET /api/user/{login}", h.GetPostByUserLogin)                      // получение постов по имени пользователя
	authMux.HandleFunc("POST /api/post/{id}", h.CreateComment)                         // добавления комментария к посту
	authMux.HandleFunc("POST /api/posts", h.CreatePost)                                // создание поста
	authMux.HandleFunc("GET /api/post/{post_id}/upvote", h.Upvote)                     // лайк
	authMux.HandleFunc("GET /api/post/{post_id}/downvote", h.Downvote)                 // дизлайк
	authMux.HandleFunc("GET /api/post/{post_id}/unvote", h.Unvote)                     // убрать голос
	authMux.HandleFunc("DELETE /api/post/{id}", h.DeletePostByID)                      // удалить пост по id
	authMux.HandleFunc("DELETE /api/post/{post_id}/{comment_id}", h.DeleteCommentByID) // удаление комментария у поста

	mux.Handle("/api/", middleware.Auth(authMux))

	return mux
}
