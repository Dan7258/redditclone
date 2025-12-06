package handler

import (
	"encoding/json"
	"net/http"
	"redditclone/internal/models"
	"redditclone/pkg/hash"
	"redditclone/pkg/jwt"
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
	user.Password, err = hash.HashPassword(user.Password)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.Logger.Debug("Registering user", user)
	err = h.db.CreateUser(user)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		h.Logger.Error("Registering user", user)
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
	userInDB, err := h.db.GetUserWithPasswordByUsername(user.Username)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "User not found")
		return
	}
	if !hash.CmpPasswordAndHash(user.Password, userInDB.Password) {
		jsonError(w, http.StatusUnauthorized, "Incorrect password")
		return
	}
	token, err := jwt.GenerateJWT(userInDB)
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
