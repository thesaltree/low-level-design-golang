package main

type NotificationType string

const (
	FriendRequestNotificationType         NotificationType = "FriendRequest"
	FriendRequestAcceptedNotificationType NotificationType = "FriendRequestAccepted"
	CommentNotificationType               NotificationType = "Comment"
	LikeNotificationType                  NotificationType = "Like"
	MentionNotificationType               NotificationType = "Mention"
	MessageNotificationType               NotificationType = "Message"
)

type Notification struct {
	ID      string
	Type    NotificationType
	Content string
	UserID  int
}

func NewNotification(id string, notifType NotificationType, content string, userID int) *Notification {
	return &Notification{ID: id, Type: notifType, Content: content, UserID: userID}
}
