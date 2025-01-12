package main

import "fmt"

type IUser interface {
	notification()
}

type User struct {
	id    int
	name  string
	email string
}

func (u *User) notification() {
	// Send email to user
	fmt.Printf("Email sent to %s\n", u.email)
}
