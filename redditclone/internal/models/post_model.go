package models

import (
	"sync"
	"time"
)

type PostMemory struct {
	mu      sync.Mutex
	Storage map[int]*Post
	counter int
}

type Post struct {
	Id       int       `json:"id,omitempty"`
	Score    int       `json:"score,omitempty"`
	Author   *User     `json:"author,omitempty"`
	Created  time.Time `json:"created,omitempty"`
	Category string    `json:"category,omitempty"`
	Title    string    `json:"title,omitempty"`
	Type     string    `json:"type,omitempty"`
	Text     string    `json:"text,omitempty"`
}

func NewPostMemory() *PostMemory {
	return &PostMemory{
		Storage: make(map[int]*Post),
		counter: 1,
		mu:      sync.Mutex{},
	}
}

func (d *PostMemory) AddPost(post *Post) (*Post, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	post.Id = d.counter
	post.Created = time.Now()
	d.Storage[post.Id] = post
	d.counter++
	return post, nil
}

func (d *PostMemory) RemovePost(post *Post) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.Storage, post.Id)
}

func (d *PostMemory) GetAllPosts() []*Post {
	d.mu.Lock()
	defer d.mu.Unlock()
	posts := make([]*Post, 0, len(d.Storage))
	for _, post := range d.Storage {
		posts = append(posts, post)
	}
	return posts
}

func (d *PostMemory) GetPostsByCategory(category string) []*Post {
	d.mu.Lock()
	defer d.mu.Unlock()
	posts := make([]*Post, 0)
	for _, post := range d.Storage {
		if post.Category == category {
			posts = append(posts, post)
		}
	}
	return posts
}

func (d *PostMemory) GetPostById(id int) *Post {
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, post := range d.Storage {
		if post.Id == id {
			return post
		}
	}
	return nil
}

func (d *PostMemory) DeletePostById(id int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.Storage, id)
}
