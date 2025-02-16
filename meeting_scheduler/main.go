package main

import (
	"fmt"
	"time"
)

// user code flow
func main() {

	// prerequisites before scheduling the meeting

	// create rooms1
	room1 := &MeetingRoom{id: 1, capacity: 10, name: "Room 1", location: location{}, calendar: &Calendar{
		interval: make(map[int]*interval),
	}}

	// create rooms2
	room2 := &MeetingRoom{id: 2, capacity: 10, name: "Room 2", location: location{}, calendar: &Calendar{
		interval: make(map[int]*interval),
	}}

	// create user
	users := []*User{
		&User{id: 1, name: "sam", email: "sam@email.com"},
		&User{id: 2, name: "ron", email: "ron@email.com"},
		&User{id: 3, name: "don", email: "don@email.com"},
		&User{id: 4, name: "him", email: "him@email.com"},
	}

	// first we will have already present rooms
	scheduler := NewMeetingScheduler()

	scheduler.rooms = append(scheduler.rooms, []*MeetingRoom{room1, room2}...)

	interval1 := &interval{
		id:    1,
		date:  time.Date(2009, 11, 17, 0, 0, 0, 0, time.UTC),
		start: time.Date(2009, 11, 17, 1, 0, 0, 0, time.UTC),
		end:   time.Date(2009, 11, 17, 1, 30, 0, 0, time.UTC)}

	// prerequisites complete,Now we start the scheduling meeting, taking sam as host and ron as participant
	_, err := scheduler.BookMeeting(1, "Daily status", []*User{users[1], users[2]}, *users[0], interval1, 10)
	if err != nil {
		fmt.Printf("\nError while scheduling meeting:%#v", err)
	}

	// lets try to book the same room again
	_, err = scheduler.BookMeeting(2, "Party", []*User{users[1], users[2]}, *users[0], interval1, 10)
	if err != nil {
		fmt.Printf("\nError while scheduling meeting:%#v", err)
	}

	// lets try to book room with more capacity
	_, err = scheduler.BookMeeting(3, "Discussion", []*User{users[1], users[2]}, *users[0], interval1, 15)
	if err != nil {
		fmt.Printf("\nError while scheduling meeting:%#v", err)
	}

	// book for other interval
	interval2 := &interval{
		id:    2,
		date:  time.Date(2009, 11, 17, 0, 0, 0, 0, time.UTC),
		start: time.Date(2009, 11, 17, 2, 0, 0, 0, time.UTC),
		end:   time.Date(2009, 11, 17, 2, 30, 0, 0, time.UTC)}

	// lets try to book room with more capacity
	meetinIDx, err := scheduler.BookMeeting(1, "Discussion", []*User{users[1], users[3]}, *users[0], interval2, 3)
	if err != nil {
		fmt.Printf("\nError while scheduling meeting:%#v", err)
	}

	// lets try to book room with more capacity
	_, err = scheduler.BookMeeting(1, "Discussion", []*User{users[1], users[3]}, *users[0], interval2, 3)
	if err != nil {
		fmt.Printf("\nError while scheduling meeting:%#v", err)
	}

	// cancel the meeting and book again
	scheduler.CancelMeeting(meetinIDx)

	// lets try to book room with more capacity
	_, err = scheduler.BookMeeting(1, "Discussion again", []*User{users[1], users[3]}, *users[0], interval2, 3)
	if err != nil {
		fmt.Printf("\nError while scheduling meeting:%#v", err)
	}
}
