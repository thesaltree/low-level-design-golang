package main

import (
	"fmt"
	"sync"
)

var (
	userManagerInstance *UserManager
	userOnce            sync.Once
)

type UserManager struct {
	users map[int]*User
	mu    sync.RWMutex
}

func GetUserManagerInstance() *UserManager {
	userOnce.Do(func() {
		userManagerInstance = &UserManager{users: make(map[int]*User)}
	})
	return userManagerInstance
}

func (um *UserManager) GetUserByID(userID int) (*User, error) {
	um.mu.RLock()
	defer um.mu.RUnlock()
	user, ok := um.users[userID]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (um *UserManager) AddUser(user *User) {
	um.mu.Lock()
	defer um.mu.Unlock()
	um.users[user.ID] = user
	fmt.Printf("User added: %d\n", user.ID)
}

func (um *UserManager) RemoveUser(userID int) {
	um.mu.Lock()
	defer um.mu.Unlock()
	delete(um.users, userID)
	fmt.Printf("User removed: %d\n", userID)
}

func (um *UserManager) LoginUser(email, password string) (*User, error) {
	um.mu.RLock()
	defer um.mu.RUnlock()
	for _, u := range um.users {
		if u.Email == email && u.Password == password {
			fmt.Printf("Login successful for email: %s\n", email)
			return u, nil
		}
	}
	return nil, fmt.Errorf("invalid email or password")
}

func (um *UserManager) UpdateUser(user *User) error {
	um.mu.Lock()
	defer um.mu.Unlock()
	um.users[user.ID] = user
	fmt.Printf("User updated: %d\n", user.ID)
	return nil
}

func (um *UserManager) AddFriend(requesterID, receiverID int) error {
	requester, err := um.GetUserByID(requesterID)
	if err != nil {
		return err
	}

	receiver, err := um.GetUserByID(receiverID)
	if err != nil {
		return err
	}

	requester.AddFriend(receiver)
	receiver.AddFriend(requester)
	fmt.Printf("Friendship added between users: %d and %d\n", requesterID, receiverID)
	return nil
}
