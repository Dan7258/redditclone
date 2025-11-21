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
			return db.Preload("Author", func(db *gorm.DB) *gorm.DB {
				return db.Omit("password")
			})
		}).
		Preload("Votes", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User", func(db *gorm.DB) *gorm.DB {
				return db.Omit("password")
			})
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
			return db.Preload("Author", func(db *gorm.DB) *gorm.DB {
				return db.Omit("password")
			})
		}).
		Preload("Votes", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User", func(db *gorm.DB) *gorm.DB {
				return db.Omit("password")
			})
		}).
		Where("category = ?", category).Find(&posts)
	return posts
}

func (db *PostgresDB) GetPostsByUsername(username string) []models.Post {
	var posts []models.Post
	var user models.User
	if err := db.Conn.Where("username = ?", username).Select("id").First(&user).Error; err != nil {
		return []models.Post{}
	}
	db.Conn.
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		Where("author_id = ?", user.ID).
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Author", func(db *gorm.DB) *gorm.DB {
				return db.Omit("password")
			})
		}).
		Preload("Votes", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User", func(db *gorm.DB) *gorm.DB {
				return db.Omit("password")
			})
		}).
		Find(&posts)
	return posts
}

func (db *PostgresDB) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	err := db.Conn.
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Author", func(db *gorm.DB) *gorm.DB {
				return db.Omit("password")
			})
		}).
		Preload("Votes", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User", func(db *gorm.DB) *gorm.DB {
				return db.Omit("password")
			})
		}).
		First(&post, id).Error
	return &post, err
}

func (db *PostgresDB) UpdatePost(post *models.Post) error {
	return db.Conn.Save(post).Error
}

func (db *PostgresDB) DeletePostByID(id uint) error {
	return db.Conn.Delete(&models.Post{}, id).Error
}
