package main

import (
	"fmt"
	"sync"
)

type BookingStatus string

const (
	BOOKED BookingStatus = "booked"
	CANCEL BookingStatus = "cancel"
)

type Meeting struct {
	id                int
	name              string
	room              *MeetingRoom
	participant       []*User
	host              *User
	meetingDurationId int
	status            BookingStatus
	mu                sync.Mutex
}

func NewMeeting(id int, name string, meetingDurationId int, room *MeetingRoom, host *User) *Meeting {
	return &Meeting{
		id:                id,
		name:              name,
		room:              room,
		meetingDurationId: meetingDurationId,
		participant:       []*User{},
		host:              host,
		status:            BOOKED,
	}
}

func (m *Meeting) AddParticipant(u ...*User) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.participant = append(m.participant, u...)
}

func (m *Meeting) RemoveParticipant(userId int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// remove user from the participant list
	for i, user := range m.participant {
		if user.id == userId {
			m.participant = append(m.participant[:i], m.participant[i+1:]...)
		}
	}
}

func (m *Meeting) setStatus(status BookingStatus) {
	m.status = status
}

func (m *Meeting) CancelMeeting() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.setStatus(CANCEL)
	m.room.CancelRoom(m.meetingDurationId)
	m.notifyParticipants(CANCEL)
}

// notify use oberver pattern to notify all the participants
func (m *Meeting) notifyParticipants(status BookingStatus) {
	fmt.Printf("\nMeeting %s has been %s\n", m.name, status)
	for _, user := range m.participant {
		user.notification()
	}
}
