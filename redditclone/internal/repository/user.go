package repository

import "redditclone/internal/models"

func (db *PostgresDB) CreateUser(user *models.User) error {
	return db.Conn.Create(user).Error
}

func (db *PostgresDB) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := db.Conn.Omit("password").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *PostgresDB) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.Conn.Omit("password").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *PostgresDB) GetUserWithPasswordByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.Conn.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
