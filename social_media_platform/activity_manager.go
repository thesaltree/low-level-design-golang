package main

import (
	"fmt"
)

type ActivityFacade struct {
	UserManager         *UserManager
	PostManager         *PostManager
	NotificationManager *NotificationManager
}

func NewActivityFacade() *ActivityFacade {
	return &ActivityFacade{
		UserManager:         GetUserManagerInstance(),
		PostManager:         GetPostManagerInstance(),
		NotificationManager: GetNotificationManagerInstance(),
	}
}

// User related operations

func (af *ActivityFacade) AddUser(user *User) {
	af.UserManager.AddUser(user)
}

func (af *ActivityFacade) RemoveUser(userID int) {
	af.UserManager.RemoveUser(userID)
}

func (af *ActivityFacade) LoginUser(email, password string) (*User, error) {
	return af.UserManager.LoginUser(email, password)
}

func (af *ActivityFacade) SendFriendRequest(requesterID, receiverID int) error {
	_, err := af.UserManager.GetUserByID(receiverID)
	if err != nil {
		return err
	}

	af.NotificationManager.AddNotification(receiverID, FriendRequestNotificationType, fmt.Sprintf("%d has sent you a friend request", requesterID))

	fmt.Printf("Friend request sent to user %d\n", receiverID)
	return nil
}

func (af *ActivityFacade) AcceptFriendRequest(requesterID, receiverID int) error {
	err := af.UserManager.AddFriend(requesterID, receiverID)
	if err != nil {
		return err
	}

	af.NotificationManager.AddNotification(requesterID, FriendRequestAcceptedNotificationType, fmt.Sprintf("%d has accepted your friend request", receiverID))

	fmt.Printf("Friend request accepted by user %d\n", receiverID)
	return nil
}

// Post related operations

func (af *ActivityFacade) AddPost(post *Post) error {
	user, err := af.UserManager.GetUserByID(post.UserID)
	if err != nil {
		return err
	}

	af.PostManager.AddPost(post, user)
	fmt.Printf("Post added: %d\n", post.ID)
	return nil
}

func (af *ActivityFacade) GetFeedPosts(userID int) ([]*Post, error) {

	user, err := af.UserManager.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return af.PostManager.GetUserFeed(user)
}

func (af *ActivityFacade) LikePost(userID, postID int) error {
	post, err := af.PostManager.LikePost(postID)
	if err != nil {
		return err
	}

	af.NotificationManager.AddNotification(post.UserID, LikeNotificationType, fmt.Sprintf("%d has liked your post: %d", userID, post.ID))

	fmt.Printf("Post liked: %d\n", post.ID)
	return nil
}

func (af *ActivityFacade) CommentPost(userID, postID int, content string) error {
	user, err := af.UserManager.GetUserByID(userID)
	if err != nil {
		return err
	}

	post, err := af.PostManager.CommentPost(user, postID, content)

	af.NotificationManager.AddNotification(post.UserID, CommentNotificationType, fmt.Sprintf("%d has commented on your post: %d", userID, post.ID))

	fmt.Printf("Comment added to post: %d\n", post.ID)
	return nil
}

func (af *ActivityFacade) MentionUserInPost(postID, mentionedUserID int) error {
	post, err := af.PostManager.GetPost(postID)
	if err != nil {
		return err
	}

	af.NotificationManager.AddNotification(post.UserID, MentionNotificationType, fmt.Sprintf("%d has mentioned you in their post: %d", mentionedUserID, post.ID))
	return nil
}

func (af *ActivityFacade) UnpublishPost(postID int) error {
	return af.PostManager.UnpublishPost(postID)
}

func (af *ActivityFacade) PublishPost(postID int) error {
	return af.PostManager.PublishPost(postID)
}

func (af *ActivityFacade) UpdatePost(postID int, content string) error {
	return af.PostManager.UpdatePost(postID, content)
}

func (af *ActivityFacade) HidePostFromUser(postID int, userID int) error {
	_, err := af.UserManager.GetUserByID(userID)
	if err != nil {
		return err
	}

	return af.PostManager.HidePostFromUser(postID, userID)
}

func (af *ActivityFacade) UnhidePostFromUser(postID int, userID int) error {
	_, err := af.UserManager.GetUserByID(userID)
	if err != nil {
		return err
	}

	return af.PostManager.UnhidePostFromUser(postID, userID)
}

func (af *ActivityFacade) EnableComments(postID int) error {
	return af.PostManager.EnableComments(postID)
}

func (af *ActivityFacade) DisableComments(postID int) error {
	return af.PostManager.DisableComments(postID)
}

// Notification related operations

func (af *ActivityFacade) GetNotifications(userID int) ([]*Notification, error) {
	return af.NotificationManager.GetNotificationsForUser(userID)
}
