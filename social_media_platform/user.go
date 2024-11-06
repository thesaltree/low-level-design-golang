package main

type User struct {
	ID             int
	Name           string
	Email          string
	Password       string
	DisplayPicture *string
	Bio            *string
	friends        map[int]*User
	posts          []*Post
}

func NewUser(id int, name, email, password, displayPicture, bio string) *User {
	return &User{ID: id, Name: name, Email: email, Password: password, DisplayPicture: &displayPicture, Bio: &bio, friends: make(map[int]*User), posts: make([]*Post, 0)}
}

func (u *User) AddPost(post *Post) {
	u.posts = append(u.posts, post)
}

func (u *User) GetPosts() []*Post {
	return u.posts
}

func (u *User) AddFriend(friend *User) {
	u.friends[friend.ID] = friend
}

func (u *User) RemoveFriend(friend *User) {
	delete(u.friends, friend.ID)
}

func (u *User) GetFriends() []*User {
	friends := make([]*User, 0, len(u.friends))
	for _, friend := range u.friends {
		friends = append(friends, friend)
	}
	return friends
}
