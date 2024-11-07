package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	notificationManagerInstance *NotificationManager
	notifOnce                   sync.Once
)

type NotificationManager struct {
	notifications map[int][]*Notification
	mu            sync.RWMutex
}

func GetNotificationManagerInstance() *NotificationManager {
	notifOnce.Do(func() {
		notificationManagerInstance = &NotificationManager{notifications: make(map[int][]*Notification)}
	})
	return notificationManagerInstance
}

func (nm *NotificationManager) AddNotification(userID int, notificationType NotificationType, message string) {
	nm.mu.Lock()
	defer nm.mu.Unlock()

	notification := NewNotification(fmt.Sprintf("notification-%d", time.Now().UnixMicro()), notificationType, message, userID)
	nm.notifications[userID] = append(nm.notifications[userID], notification)
}

func (nm *NotificationManager) GetNotificationsForUser(userID int) ([]*Notification, error) {
	nm.mu.RLock()
	defer nm.mu.RUnlock()

	notifications, ok := nm.notifications[userID]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return notifications, nil
}
