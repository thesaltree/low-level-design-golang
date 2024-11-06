package main

import (
	"fmt"
	"sync"
	"time"
)

type ActivityManager struct {
	users         map[int]*User
	posts         map[int]*Post
	notifications map[int][]*Notification
	mu            sync.RWMutex
}

var (
	activityManagerInstance *ActivityManager
	once                    sync.Once
)

func GetActivityManagerInstance() *ActivityManager {
	once.Do(func() {
		activityManagerInstance = &ActivityManager{users: make(map[int]*User), posts: make(map[int]*Post), notifications: make(map[int][]*Notification)}
	})
	return activityManagerInstance
}

// User related operations

func (am *ActivityManager) AddUser(user *User) {
	am.mu.Lock()
	defer am.mu.Unlock()
	am.users[user.ID] = user
}

func (am *ActivityManager) RemoveUser(userID int) {
	am.mu.Lock()
	defer am.mu.Unlock()
	delete(am.users, userID)
}

func (am *ActivityManager) LoginUser(email, password string) (*User, error) {
	am.mu.RLock()
	defer am.mu.RUnlock()
	for _, u := range am.users {
		if u.Email == email && u.Password == password {
			return u, nil
		}
	}
	return nil, fmt.Errorf("invalid email or password")
}

func (am *ActivityManager) SendFriendRequest(requesterID, receiverID int) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	_, exists := am.users[receiverID]
	if !exists {
		return fmt.Errorf("receiver not found")
	}

	notification := NewNotification(fmt.Sprintf("notification-%d", time.Now().UnixMicro()), FriendRequestNotificationType, fmt.Sprintf("%d has sent you a friend request", requesterID), receiverID)
	am.notifications[receiverID] = append(am.notifications[receiverID], notification)

	fmt.Printf("Friend request sent to user %d\n", receiverID)
	return nil
}

func (am *ActivityManager) AcceptFriendRequest(requesterID, receiverID int) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	requester, exists := am.users[requesterID]
	if !exists {
		return fmt.Errorf("requester not found")
	}

	receiver, exists := am.users[receiverID]
	if !exists {
		return fmt.Errorf("receiver not found")
	}

	requester.AddFriend(receiver)
	receiver.AddFriend(requester)

	notification := NewNotification(fmt.Sprintf("notification-%d", time.Now().UnixMicro()), FriendRequestAcceptedNotificationType, fmt.Sprintf("%d has accepted your friend request", requesterID), receiverID)
	am.notifications[requesterID] = append(am.notifications[requesterID], notification)

	fmt.Printf("Friend request accepted by user %d\n", receiverID)
	return nil
}

// Post related operations

func (am *ActivityManager) AddPost(post *Post) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	user, exists := am.users[post.UserID]
	if !exists {
		return fmt.Errorf("user not found")
	}

	am.posts[post.ID] = post
	user.AddPost(post)
	fmt.Printf("Post added: %d\n", post.ID)
	return nil
}

func (am *ActivityManager) GetFeedPosts(user *User) ([]*Post, error) {
	am.mu.RLock()
	defer am.mu.RUnlock()
	friends := user.GetFriends()
	posts := make([]*Post, 0)

	posts = append(posts, user.GetPosts()...)

	for _, friend := range friends {
		userPosts := friend.GetPosts()
		for _, post := range userPosts {
			if !post.IsHiddenFromUser(*user) && post.IsPublished {
				posts = append(posts, post)
			}
		}
	}

	return posts, nil
}

func (am *ActivityManager) LikePost(userID, postID int) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	post, exists := am.posts[postID]
	if !exists {
		return fmt.Errorf("post not found")
	}

	post.Likes++
	notification := NewNotification(fmt.Sprintf("notification-%d", time.Now().UnixMicro()), LikeNotificationType, fmt.Sprintf("%d has liked your post: %d", userID, post.ID), post.UserID)
	am.notifications[post.UserID] = append(am.notifications[post.UserID], notification)

	fmt.Printf("Post liked: %d\n", post.ID)
	return nil
}

func (am *ActivityManager) CommentPost(userID, postID int, content string) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	post, exists := am.posts[postID]
	if !exists {
		return fmt.Errorf("post not found")
	}

	comment := NewComment(fmt.Sprintf("comment-%d", time.Now().UnixNano()), postID, content, am.users[userID])
	post.Comments = append(post.Comments, comment)
	notification := NewNotification(fmt.Sprintf("notification-%d", time.Now().UnixMicro()), CommentNotificationType, fmt.Sprintf("%d has commented on your post: %d", userID, post.ID), post.UserID)
	am.notifications[post.UserID] = append(am.notifications[post.UserID], notification)

	fmt.Printf("Comment added to post: %d\n", post.ID)
	return nil
}

func (am *ActivityManager) MentionUserInPost(postID, mentionedUserID int) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	post, exists := am.posts[postID]
	if !exists {
		return fmt.Errorf("post not found")
	}

	notification := NewNotification(fmt.Sprintf("notification-%d", time.Now().UnixMicro()), MentionNotificationType, fmt.Sprintf("%d has mentioned you in their post: %d", mentionedUserID, post.ID), post.UserID)
	am.notifications[mentionedUserID] = append(am.notifications[mentionedUserID], notification)
	return nil
}

func (am *ActivityManager) UnpublishPost(postID int) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	post, exists := am.posts[postID]
	if !exists {
		return fmt.Errorf("post not found")
	}

	post.UnpublishPost()
	return nil
}

func (am *ActivityManager) PublishPost(postID int) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	post, exists := am.posts[postID]
	if !exists {
		return fmt.Errorf("post not found")
	}

	post.PublishPost()
	return nil
}

func (am *ActivityManager) UpdatePost(postID int, content string) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	post, exists := am.posts[postID]
	if !exists {
		return fmt.Errorf("post not found")
	}

	post.UpdateContent(content)
	return nil
}

func (am *ActivityManager) HidePostFromUser(postID int, userID int) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	post, exists := am.posts[postID]
	if !exists {
		return fmt.Errorf("post not found")
	}

	user := am.users[userID]

	post.HideFromUser(*user)
	post.HiddenFromUsers[userID] = user
	return nil
}

func (am *ActivityManager) UnhidePostFromUser(postID int, userID int) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	post, exists := am.posts[postID]
	if !exists {
		return fmt.Errorf("post not found")
	}

	post.UnhideFromUser(*am.users[userID])
	delete(post.HiddenFromUsers, userID)
	return nil
}

func (am *ActivityManager) DisableComments(postID int) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	post, exists := am.posts[postID]
	if !exists {
		return fmt.Errorf("post not found")
	}

	post.DisableComments()
	return nil
}

func (am *ActivityManager) EnableComments(postID int) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	post, exists := am.posts[postID]
	if !exists {
		return fmt.Errorf("post not found")
	}

	post.EnableComments()
	return nil
}

// Notification related operations

func (am *ActivityManager) GetNotifications(userID int) ([]*Notification, error) {
	am.mu.RLock()
	defer am.mu.RUnlock()

	notifications, exists := am.notifications[userID]

	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	return notifications, nil
}
