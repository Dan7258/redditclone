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

	CreateComment(comment *Comment) error
}
