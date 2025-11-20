package models

import "time"

type Comment struct {
	ID      uint      `gorm:"primary_key;auto_increment" json:"id"`
	PostID  uint      `gorm:"primary_key" json:"post_id"`
	Author  User      `json:"author"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}
