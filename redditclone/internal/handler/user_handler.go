package handler

import (
	"encoding/json"
	"net/http"
	"redditclone/internal/models"
	"redditclone/jwt"
)

type jwtResponse struct {
	Token string `json:"token"`
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = h.db.CreateUser(user)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := jwt.GenerateJWT(user)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	resp := jwtResponse{
		Token: token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		jsonError(w, http.StatusBadRequest, err.Error())
		return
	}
	userInMemory, err := h.db.GetUserByUsername(user.Username)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "User not found")
		return
	}
	token, err := jwt.GenerateJWT(userInMemory)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	resp := jwtResponse{
		Token: token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
