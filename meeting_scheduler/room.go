package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type location struct {
	building string
	floor    string
}

type MeetingRoom struct {
	id       int
	capacity int
	name     string
	location location
	calender *Calendar
}

func NewMeetingRoom(id int, capacity int, name string, location location) *MeetingRoom {
	return &MeetingRoom{
		id:       id,
		capacity: capacity,
		name:     name,
		location: location,
		calender: &Calendar{},
	}
}

func (m *MeetingRoom) GetName() string {
	return m.name
}

func (m *MeetingRoom) BookRoom(capacity int, dur interval) error {

	// Check if the room has the capacity
	if !m.hasCapacity(capacity) {
		return errors.New("Room does not have the capacity")
	}

	// Check if the room calender is free for the given interval
	if m.calender.isFree(dur) {
		fmt.Printf("\nRoom %s is booked for the given interval startTime:%s endTime:%v\n", m.name, dur.start.String(), dur.end.String())
		m.calender.bookInterval(dur)
		return nil
	}

	return errors.New("Room is not free for the given interval")
}

func (m *MeetingRoom) hasCapacity(capacity int) bool {
	// Check if the room has the capacity
	if m.capacity >= capacity {
		return true
	}
	return false
}
