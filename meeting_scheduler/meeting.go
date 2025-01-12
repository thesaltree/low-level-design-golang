package main

import "fmt"

type Meeting struct {
	id          int
	name        string
	room        *MeetingRoom
	participant []*User
	host        *User
}

func NewMeeting(id int, name string, room *MeetingRoom, host *User) *Meeting {
	return &Meeting{
		id:          id,
		name:        name,
		room:        room,
		participant: []*User{},
		host:        host,
	}
}

func (m *Meeting) addParticipant(u ...*User) {
	m.participant = append(m.participant, u...)
}

func (m *Meeting) removeParticipant(userId int) {
	// remove user from the participant list
	for i, user := range m.participant {
		if user.id == userId {
			m.participant = append(m.participant[:i], m.participant[i+1:]...)
		}
	}
}

// notify use oberver pattern to notify all the participants
func (m *Meeting) notifyParticipants() {
	fmt.Printf("\nMeeting %s has been scheduled\n", m.name)
	for _, user := range m.participant {
		user.notification()
	}
}
