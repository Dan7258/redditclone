package models

import "time"

type Comment struct {
	ID       uint      `gorm:"primary_key;auto_increment" json:"id"`
	PostID   uint      `json:"post_id"`
	AuthorID uint      `gorm:"not null" json:"-"`
	Author   User      `gorm:"foreignkey:AuthorID" json:"author"`
	Body     string    `json:"body"`
	Created  time.Time `json:"created"`
}
