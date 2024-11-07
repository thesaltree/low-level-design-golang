package main

import (
	"sync"
	"time"
)

type Post struct {
	ID              int
	UserID          int
	Content         string
	IsPublished     bool
	URLs            []*string
	Likes           int
	Comments        []*Comment
	PublishedAt     time.Time
	CommentsEnabled bool
	HiddenFromUsers map[int]bool
	mu              sync.RWMutex
}

func NewPost(id int, userID int, content string, urls []*string) *Post {
	return &Post{
		ID:              id,
		UserID:          userID,
		Content:         content,
		IsPublished:     true,
		URLs:            urls,
		Likes:           0,
		Comments:        make([]*Comment, 0),
		PublishedAt:     time.Now(),
		CommentsEnabled: true,
		HiddenFromUsers: make(map[int]bool),
	}
}

func (p *Post) AddComment(comment *Comment) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Comments = append(p.Comments, comment)
}

func (p *Post) GetComments() []*Comment {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.Comments
}

func (p *Post) Like() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Likes++
}

func (p *Post) GetLikes() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.Likes
}

func (p *Post) HideFromUser(userID int) {
	p.HiddenFromUsers[userID] = true
}

func (p *Post) UnhideFromUser(userID int) {
	delete(p.HiddenFromUsers, userID)
}

func (p *Post) IsHiddenFromUser(user User) bool {
	_, ok := p.HiddenFromUsers[user.ID]
	return ok
}

func (p *Post) EnableComments() {
	p.CommentsEnabled = true
}

func (p *Post) DisableComments() {
	p.CommentsEnabled = false
}

func (p *Post) IsCommentsEnabled() bool {
	return p.CommentsEnabled
}

func (p *Post) PublishPost() {
	p.IsPublished = true
	p.PublishedAt = time.Now()
}

func (p *Post) UnpublishPost() {
	p.IsPublished = false
}

func (p *Post) UpdateContent(content string) {
	p.Content = content
}
