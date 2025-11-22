package models

import "time"

type Vote struct {
	ID      uint      `gorm:"primary_key" json:"-"`
	UserID  uint      `json:"-"`
	User    User      `gorm:"foreignkey:UserID" json:"user"`
	PostID  uint      `json:"-"`
	Vote    int       `json:"vote"`
	Created time.Time `json:"-"`
}
