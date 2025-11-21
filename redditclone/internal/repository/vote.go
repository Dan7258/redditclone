package repository

import "redditclone/internal/models"

func (db *PostgresDB) GetVoteByUserIDAndPostID(userID, postID uint) (models.Vote, error) {
	var vote models.Vote
	err := db.Conn.Where("user_id = ? AND post_id = ?", userID, postID).First(&vote).Error
	return vote, err
}

func (db *PostgresDB) Vote(vote *models.Vote) error {
	return db.Conn.Save(vote).Error
}

func (db *PostgresDB) DeleteVoteByUserIDAndPostID(userID, postID uint) error {
	return db.Conn.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&models.Vote{}).Error
}
