package models

import "time"

type Comment struct {
	ID      uint      `gorm:"primary_key;auto_increment" json:"id"`
	PostID  uint      `json:"post_id"`
	Author  User      `gorm:"foreignkey:ID" json:"author"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}
