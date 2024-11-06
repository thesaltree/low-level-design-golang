package main

type Comment struct {
	ID      string
	PostID  int
	Content string
	User    *User
}

func NewComment(id string, postID int, content string, user *User) *Comment {
	return &Comment{ID: id, PostID: postID, Content: content, User: user}
}
