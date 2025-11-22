package repository

import "redditclone/internal/models"

func (db *PostgresDB) CreateComment(comment *models.Comment) error {
	return db.Conn.Create(comment).Error
}

func (db *PostgresDB) DeleteCommentByID(id uint) error {
	return db.Conn.Delete(models.Comment{}, id).Error
}
