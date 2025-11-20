package repository

import "redditclone/internal/models"

func (db *PostgresDB) CreateComment(comment *models.Comment) error {
	return db.Conn.Create(comment).Error
}
