package models

import (
	"gorm.io/gorm"
)

type Model interface {
	GetConn() *gorm.DB
	ConnectToDatabase() error
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)

	CreatePost(post *Post) error
	GetAllPosts() []Post
	GetPostsByCategory(category string) []Post
	GetPostByID(id uint) (*Post, error)
	DeletePostByID(id uint) error
	UpdatePost(post *Post) error
	GetPostsByUsername(username string) []Post

	CreateComment(comment *Comment) error
	DeleteCommentByID(id uint) error

	GetVoteByUserIDAndPostID(userID, postID uint) (Vote, error)
	Vote(vote *Vote) error
	DeleteVoteByUserIDAndPostID(userID, postID uint) error
}
