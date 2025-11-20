package repository

import (
	"gorm.io/gorm"
	"redditclone/internal/models"
)

func (db *PostgresDB) CreatePost(post *models.Post) error {
	return db.Conn.Create(post).Error
}

func (db *PostgresDB) GetAllPosts() []models.Post {
	var posts []models.Post
	db.Conn.
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		Find(&posts)
	return posts
}

func (db *PostgresDB) GetPostsByCategory(category string) []models.Post {
	var posts []models.Post
	db.Conn.
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		Where("category = ?", category).Find(&posts)
	return posts
}

func (db *PostgresDB) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	err := db.Conn.
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		First(&post, id).Error
	return &post, err
}

func (db *PostgresDB) DeletePostByID(id uint) error {
	return db.Conn.Delete(&models.Post{}, id).Error
}
