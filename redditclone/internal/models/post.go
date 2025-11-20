package models

import "time"

type Post struct {
	ID               uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title            string    `json:"title"`
	Author           User      `json:"author"`
	Category         string    `json:"category"`
	Score            int       `json:"score"`
	Views            int       `json:"views"`
	Type             string    `json:"type"`
	Text             string    `json:"text"`
	UpvotePercentage int       `json:"upvote_percentage"`
	Created          time.Time `json:"created"`
}
