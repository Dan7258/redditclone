package models

type User struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `gorm:"not null" json:"password,omitempty"`
}
