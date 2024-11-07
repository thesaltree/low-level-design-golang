package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	postManagerInstance *PostManager
	postOnce            sync.Once
)

type PostManager struct {
	posts map[int]*Post
	mu    sync.RWMutex
}

func GetPostManagerInstance() *PostManager {
	postOnce.Do(func() {
		postManagerInstance = &PostManager{posts: make(map[int]*Post)}
	})
	return postManagerInstance
}

func (pm *PostManager) GetPost(postID int) (*Post, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	post, exists := pm.posts[postID]
	if !exists {
		return nil, fmt.Errorf("post not found")
	}
	return post, nil
}

func (pm *PostManager) AddPost(post *Post, user *User) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.posts[post.ID] = post

	user.AddPost(post)
}

func (pm *PostManager) LikePost(postID int) (*Post, error) {
	pm.mu.Lock()
	post := pm.posts[postID]
	pm.mu.Unlock()

	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()

	post.Like()
	return post, nil
}

func (pm *PostManager) CommentPost(user *User, postID int, content string) (*Post, error) {
	pm.mu.Lock()
	post := pm.posts[postID]
	pm.mu.Unlock()

	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()

	comment := NewComment(fmt.Sprintf("comment-%d", time.Now().UnixNano()), postID, content, user)
	post.AddComment(comment)

	return post, nil
}

func (pm *PostManager) GetUserFeed(user *User) ([]*Post, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	var feed []*Post

	friends := user.GetFriends()

	feed = append(feed, user.GetPosts()...)

	for _, friend := range friends {
		userPosts := friend.GetPosts()
		for _, post := range userPosts {
			if !post.IsHiddenFromUser(*user) && post.IsPublished {
				feed = append(feed, post)
			}
		}
	}

	return feed, nil
}

func (pm *PostManager) UnpublishPost(postID int) error {
	pm.mu.Lock()
	post := pm.posts[postID]
	pm.mu.Unlock()

	if post == nil {
		return fmt.Errorf("post not found")
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()

	post.UnpublishPost()
	fmt.Printf("Post %d unpublished\n", postID)
	return nil
}

func (pm *PostManager) PublishPost(postID int) error {
	pm.mu.Lock()
	post := pm.posts[postID]
	pm.mu.Unlock()

	if post == nil {
		return fmt.Errorf("post not found")
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()

	post.PublishPost()
	fmt.Printf("Post %d published\n", postID)
	return nil
}

func (pm *PostManager) UpdatePost(postID int, content string) error {
	pm.mu.Lock()
	post := pm.posts[postID]
	pm.mu.Unlock()

	if post == nil {
		return fmt.Errorf("post not found")
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()

	post.UpdateContent(content)
	fmt.Printf("Post %d updated with content: %s\n", postID, content)
	return nil
}

func (pm *PostManager) HidePostFromUser(postID, userID int) error {
	pm.mu.Lock()
	post := pm.posts[postID]
	pm.mu.Unlock()

	if post == nil {
		return fmt.Errorf("post not found")
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()

	post.HideFromUser(userID)
	post.HiddenFromUsers[userID] = true
	fmt.Printf("Post %d hidden from user %d\n", postID, userID)
	return nil
}

func (pm *PostManager) UnhidePostFromUser(postID, userID int) error {
	pm.mu.Lock()
	post := pm.posts[postID]
	pm.mu.Unlock()

	if post == nil {
		return fmt.Errorf("post not found")
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()

	post.UnhideFromUser(userID)
	delete(post.HiddenFromUsers, userID)
	fmt.Printf("Post %d unhidden from user %d\n", postID, userID)
	return nil
}

func (pm *PostManager) EnableComments(postID int) error {
	pm.mu.Lock()
	post := pm.posts[postID]
	pm.mu.Unlock()

	if post == nil {
		return fmt.Errorf("post not found")
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()

	post.EnableComments()
	fmt.Printf("Comments enabled for post %d\n", postID)
	return nil
}

func (pm *PostManager) DisableComments(postID int) error {
	pm.mu.Lock()
	post := pm.posts[postID]
	pm.mu.Unlock()

	if post == nil {
		return fmt.Errorf("post not found")
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()

	post.DisableComments()
	fmt.Printf("Comments disabled for post %d\n", postID)
	return nil
}
