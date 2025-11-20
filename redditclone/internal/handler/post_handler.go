package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"redditclone/internal/models"
	"redditclone/jwt"
	"strconv"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	post := new(models.Post)
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	claims, ok := r.Context().Value("user").(jwt.Claims)
	if !ok {
		jsonError(w, http.StatusUnauthorized, "user not authorized")
		return
	}
	post.AuthorID = claims.ID
	err = h.db.CreatePost(post)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	resp, err := json.Marshal(post)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := h.db.GetAllPosts()
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) GetPostById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		jsonError(w, http.StatusBadRequest, "category not provided")
		return
	}
	id, err := strconv.Atoi(idStr)
	post, err := h.db.GetPostByID(uint(id))
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) GetPostsByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("category")
	if category == "" {
		jsonError(w, http.StatusBadRequest, "category not provided")
		return
	}
	posts := h.db.GetPostsByCategory(category)
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) DeletePostById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil || idStr == "" {
		jsonError(w, http.StatusBadRequest, "id not provided")
		return
	}
	claims, ok := r.Context().Value("user").(jwt.Claims)
	if !ok {
		jsonError(w, http.StatusUnauthorized, "user not authorized")
		return
	}
	post, err := h.db.GetPostByID(claims.ID)
	if post != nil {
		jsonError(w, http.StatusInternalServerError, fmt.Sprintf("post with id: %d not exists", id))
		return
	}
	if post.AuthorID != claims.ID {
		jsonError(w, http.StatusForbidden, "Forbidden")
		return
	}
	err = h.db.DeletePostByID(uint(id))
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "success"}`))
}

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	comment := new(models.Comment)
	message := struct {
		Comment string `json:"comment"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}
	idStr := r.PathValue("id")
	if idStr == "" {
		jsonError(w, http.StatusBadRequest, "id required")
		return
	}
	postId, err := strconv.Atoi(idStr)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "id required")
		return
	}
	claims, ok := r.Context().Value("user").(jwt.Claims)
	if !ok {
		jsonError(w, http.StatusUnauthorized, "user not authorized")
		return
	}
	comment.AuthorID = claims.ID
	comment.PostID = uint(postId)
	comment.Body = message.Comment

	err = h.db.CreateComment(comment)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	post, err := h.db.GetPostByID(uint(postId))
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
}
