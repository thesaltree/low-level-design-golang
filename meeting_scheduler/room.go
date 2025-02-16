package main

import (
	"fmt"
	"sync"

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
	status   BookingStatus
	location location
	calendar *Calendar
	mu       sync.Mutex
}

func NewMeetingRoom(id int, capacity int, name string, location location) *MeetingRoom {
	return &MeetingRoom{
		id:       id,
		capacity: capacity,
		name:     name,
		location: location,
		calendar: &Calendar{
			interval: make(map[int]*interval),
		},
	}
}

func (m *MeetingRoom) GetName() string {
	return m.name
}

func (m *MeetingRoom) BookRoom(capacity int, dur *interval) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if the room has the capacity
	if !m.hasCapacity(capacity) {
		return errors.New("Room does not have the capacity")
	}

	// Check if the room calendar is free for the given interval
	if m.isFree(dur) {
		fmt.Printf("\nRoom %s is booked for the given interval startTime:%s endTime:%v\n", m.name, dur.start.String(), dur.end.String())
		m.calendar.bookInterval(dur)
		return nil
	}

	return errors.New("Room is not free for the given interval")
}

func (m *MeetingRoom) CancelRoom(durId int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.calendar.cancelInterval(durId)
}

func (m *MeetingRoom) isFree(dur *interval) bool {
	return m.calendar.isFree(dur)
}

// IsFree is use to explose outside
func (m *MeetingRoom) IsFree(dur *interval) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.isFree(dur)
}

func (m *MeetingRoom) hasCapacity(capacity int) bool {
	// Check if the room has the capacity
	if m.capacity >= capacity {
		return true
	}
	return false
}
