package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"redditclone/internal/middleware"
	"redditclone/internal/models"
	"strconv"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	post := new(models.Post)
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}
	user, ok := middleware.GetUserFromContext(r.Context())
	if !ok {
		jsonError(w, http.StatusUnauthorized, "user not authorized")
		return
	}
	post.Author = user
	newPost, err := h.Models.PostMemory.AddPost(post)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	resp, err := json.Marshal(newPost)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := h.Models.PostMemory.GetAllPosts()
	resp, err := json.Marshal(posts)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *Handler) GetPostsByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("category")
	if category == "" {
		jsonError(w, http.StatusBadRequest, "category not provided")
		return
	}
	post := h.Models.PostMemory.GetPostsByCategory(category)
	resp, err := json.Marshal(post)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *Handler) DeletePostById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil || idStr == "" {
		jsonError(w, http.StatusBadRequest, "id not provided")
		return
	}
	user, ok := middleware.GetUserFromContext(r.Context())
	if !ok {
		jsonError(w, http.StatusUnauthorized, "user not authorized")
		return
	}
	post := h.Models.PostMemory.GetPostById(id)
	if post != nil {
		jsonError(w, http.StatusInternalServerError, fmt.Sprintf("post with id: %d not exists", id))
		return
	}
	if post.Author.Id != user.Id {
		jsonError(w, http.StatusForbidden, "Forbidden")
		return
	}
	h.Models.PostMemory.DeletePostById(id)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "success"}`))
}
