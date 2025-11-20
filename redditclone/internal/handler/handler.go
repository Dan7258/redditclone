package handler

import (
	"encoding/json"
	"net/http"
	"redditclone/internal/models"
)

type Handler struct {
	db models.Model
}

func InitHandler(model models.Model) *Handler {
	return &Handler{model}
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
