package models

import (
	"encoding/json"
	"time"
)

type Post struct {
	ID               uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title            string    `json:"title,omitempty"`
	AuthorID         uint      `gorm:"not null" json:"-"`
	Author           User      `gorm:"foreignkey:AuthorID" json:"author"`
	Category         string    `json:"category,omitempty"`
	Score            int       `json:"score,omitempty"`
	Views            int       `json:"views,omitempty"`
	Type             string    `json:"type,omitempty"`
	Text             string    `json:"text,omitempty"`
	UpvotePercentage int       `json:"upvote_percentage,omitempty"`
	Created          time.Time `json:"created,omitempty"`
	Comments         []Comment `gorm:"foreignkey:PostID" json:"comments,omitempty"`
	Votes            []Vote    `gorm:"foreignkey:PostID" json:"votes,omitempty"`
}

func (p *Post) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"id":                p.ID,
		"title":             p.Title,
		"author":            p.Author,
		"votes":             p.Votes,
		"category":          p.Category,
		"score":             p.Score,
		"views":             p.Views,
		"type":              p.Type,
		"text":              p.Text,
		"upvote_percentage": p.UpvotePercentage,
		"created":           p.Created,
		"comments":          p.Comments,
	}
	if p.Type == "link" {
		data["url"] = data["text"].(string)
		delete(data, "text")
	}
	return json.Marshal(data)
}

func (p *Post) UnmarshalJSON(data []byte) error {
	type Alias Post
	alias := struct {
		Url string `json:"url,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}
	if alias.Url != "" {
		p.Text = alias.Url
	}

	return err
}
