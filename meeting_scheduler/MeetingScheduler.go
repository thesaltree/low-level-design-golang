package main

import "errors"

type MeetingScheduler struct {
	rooms   []*MeetingRoom
	meeting []*Meeting
}

func NewMeetingScheduler() *MeetingScheduler {
	return &MeetingScheduler{
		rooms:   []*MeetingRoom{},
		meeting: []*Meeting{},
	}
}

func (ms *MeetingScheduler) getRoom(roomId int) *MeetingRoom {
	for _, room := range ms.rooms {
		if room.id == roomId {
			return room
		}
	}
	return nil
}

// book meeting for a given room
func (ms *MeetingScheduler) bookMeeting(roomId int, meetingName string, participant []*User, host User, dur interval, capacity int) error {

	room := ms.getRoom(roomId)
	if room == nil {
		return errors.New("Room not found")
	}

	err := room.BookRoom(capacity, dur)
	if err != nil {
		return err
	}

	meeting := NewMeeting(len(ms.meeting)+1, meetingName, room, &host)

	meeting.addParticipant(participant...)
	meeting.notifyParticipants()

	return nil
}
