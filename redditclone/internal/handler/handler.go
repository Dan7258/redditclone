package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"redditclone/internal/models"
)

type Handler struct {
	db     models.Model
	Logger *slog.Logger
}

func InitHandler(model models.Model) *Handler {
	return &Handler{
		db:     model,
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, nil))}
}

func jsonError(w http.ResponseWriter, status int, msg string) {
	resp, _ := json.Marshal(map[string]interface{}{
		"status":  status,
		"message": msg,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(resp)
}

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("web/html/")).ServeHTTP(w, r)
}

func (h *Handler) Static(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/static/", http.FileServer(http.Dir("web/"))).ServeHTTP(w, r)
}
