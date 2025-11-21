package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"redditclone/internal/models"
	"redditclone/pkg/jwt"
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

func (h *Handler) GetPostByID(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetPostByUserLogin(w http.ResponseWriter, r *http.Request) {
	login := r.PathValue("login")
	if login == "" {
		jsonError(w, http.StatusBadRequest, "category not provided")
		return
	}
	posts := h.db.GetPostsByUsername(login)
	err := json.NewEncoder(w).Encode(posts)
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

func (h *Handler) DeletePostByID(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) DeleteCommentByID(w http.ResponseWriter, r *http.Request) {
	postIdStr := r.PathValue("post_id")
	commentIdStr := r.PathValue("comment_id")
	if postIdStr == "" || commentIdStr == "" {
		jsonError(w, http.StatusBadRequest, "post_id or comment_id not provided")
		return
	}
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "post_id required")
		return
	}
	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "comment_id required")
		return
	}
	err = h.db.DeleteCommentByID(uint(commentId))
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

func (h *Handler) Upvote(w http.ResponseWriter, r *http.Request) {
	postIdStr := r.PathValue("post_id")
	if postIdStr == "" {
		jsonError(w, http.StatusBadRequest, "post_id not provided")
		return
	}
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "post_id required")
		return
	}
	claims := r.Context().Value("user").(jwt.Claims)
	err = h.Vote(1, claims.ID, uint(postId))
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	post, err := h.UpdatePostScore(uint(postId))
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

func (h *Handler) Downvote(w http.ResponseWriter, r *http.Request) {
	postIdStr := r.PathValue("post_id")
	if postIdStr == "" {
		jsonError(w, http.StatusBadRequest, "post_id not provided")
		return
	}
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "post_id required")
		return
	}
	claims := r.Context().Value("user").(jwt.Claims)
	err = h.Vote(-1, claims.ID, uint(postId))
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	post, err := h.UpdatePostScore(uint(postId))
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

func (h *Handler) Unvote(w http.ResponseWriter, r *http.Request) {
	postIdStr := r.PathValue("post_id")
	if postIdStr == "" {
		jsonError(w, http.StatusBadRequest, "post_id not provided")
		return
	}
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "post_id required")
		return
	}
	claims := r.Context().Value("user").(jwt.Claims)
	err = h.Vote(0, claims.ID, uint(postId))
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	post, err := h.UpdatePostScore(uint(postId))
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

func (h *Handler) Vote(numVote int, userID, postID uint) error {
	var err error
	if numVote == 0 {
		err = h.db.DeleteVoteByUserIDAndPostID(userID, postID)
	} else {
		vote := new(models.Vote)
		voteDB := new(models.Vote)
		vote.Vote = numVote
		vote.PostID = postID
		vote.UserID = userID
		*voteDB, err = h.db.GetVoteByUserIDAndPostID(vote.UserID, vote.PostID)
		if err != nil {
			err = h.db.Vote(vote)
		} else {
			voteDB.Vote = vote.Vote
			err = h.db.Vote(voteDB)
		}
	}
	return err
}

func (h *Handler) UpdatePostScore(postID uint) (*models.Post, error) {
	post, err := h.db.GetPostByID(postID)
	if err != nil {
		return nil, err
	}
	post.Score, post.UpvotePercentage = GetSumVotesAndUpvotePercentage(post.Votes)
	err = h.db.UpdatePost(post)
	return post, err
}

func GetSumVotesAndUpvotePercentage(votes []models.Vote) (int, int) {
	sum := 0
	upvote := 0

	for _, vote := range votes {
		sum += vote.Vote
		if vote.Vote == 1 {
			upvote++
		}
	}
	upvotePercentage := 0
	if len(votes) > 0 {
		upvotePercentage = 100 / len(votes) * upvote
	}
	return sum, upvotePercentage

}
