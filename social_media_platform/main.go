package main

import "fmt"

func main() {
	socialMedia := NewActivityFacade()

	// Simulate user actions

	// Create users and add them to the social media platform
	user1 := NewUser(1, "John Doe", "john.doe@example.com", "password123", "", "I'm a lovely person!")
	user2 := NewUser(2, "Jane Smith", "jane.smith@example.com", "secret123", "", "I'm a talented artist!")

	socialMedia.AddUser(user1)
	socialMedia.AddUser(user2)

	//Create posts and add them to the social media platform
	post1 := NewPost(1, 1, "Today, I went to the park and walked 1000 steps!", nil)
	post2 := NewPost(2, 2, "I painted a beautiful sunset on the wall!", nil)
	post3 := NewPost(3, 1, "I've been working on a new video game!", nil)

	err := socialMedia.AddPost(post1)
	if err != nil {
		fmt.Println("Error adding post:", err)
	}
	err = socialMedia.AddPost(post2)
	if err != nil {
		fmt.Println("Error adding post:", err)
	}
	err = socialMedia.AddPost(post3)
	if err != nil {
		fmt.Println("Error adding post:", err)
	}

	// Get user feeds
	err = getUserFeed(socialMedia, user1.ID)
	if err != nil {
		fmt.Println("Error getting user feed:", err)
	}
	err = getUserFeed(socialMedia, user2.ID)
	if err != nil {
		fmt.Println("Error getting user feed:", err)
	}

	// Send friend requests and accept them
	err = socialMedia.SendFriendRequest(1, 2)
	if err != nil {
		fmt.Println("Error sending friend request:", err)
	}
	err = socialMedia.AcceptFriendRequest(1, 2)
	if err != nil {
		fmt.Println("Error accepting friend request:", err)
	}

	err = getUserFeed(socialMedia, user1.ID)
	if err != nil {
		fmt.Println("Error getting user feed:", err)
	}

	// Publish and unpublish posts
	err = socialMedia.UnpublishPost(1)

	err = getUserFeed(socialMedia, user2.ID)
	if err != nil {
		fmt.Println("Error getting user feed:", err)
	}

	// Comment on a post
	err = socialMedia.CommentPost(1, 2, "I really like this post!")
	if err != nil {
		fmt.Println("Error commenting on post:", err)
	}

	comments := post2.GetComments()
	for _, comment := range comments {
		fmt.Printf("User %s: %s\n", comment.User.Name, comment.Content)
	}

	// Like a post
	err = socialMedia.LikePost(2, 3)
	if err != nil {
		fmt.Println("Error liking post:", err)
	}

	fmt.Printf("Post %d's likes: %d\n", post3.ID, post3.GetLikes())

	err = socialMedia.LikePost(2, 3)
	if err != nil {
		fmt.Println("Error liking post:", err)
	}

	fmt.Printf("Post %d's likes: %d\n", post3.ID, post3.GetLikes())
}

func getUserFeed(socialMedia *ActivityFacade, userID int) error {
	feed, err := socialMedia.GetFeedPosts(userID)
	if err != nil {
		fmt.Println("Error getting feed posts:", err)
	}
	for _, post := range feed {
		fmt.Printf("User %d: %s\n", post.UserID, post.Content)
	}

	return nil
}
