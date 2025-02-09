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

func (ms *MeetingScheduler) GetRoom(roomId int) *MeetingRoom {
	for _, room := range ms.rooms {
		if room.id == roomId {
			return room
		}
	}
	return nil
}

// book meeting for a given room
func (ms *MeetingScheduler) BookMeeting(roomId int, meetingName string, participant []*User, host User, dur *interval, capacity int) (int, error) {

	room := ms.GetRoom(roomId)
	if room == nil {
		return 0, errors.New("Room not found")
	}

	err := room.BookRoom(capacity, dur)
	if err != nil {
		return 0, err
	}

	meeting := NewMeeting(len(ms.meeting), meetingName, dur.id, room, &host)
	ms.meeting = append(ms.meeting, meeting)

	meeting.AddParticipant(participant...)

	return len(ms.meeting) - 1, nil
}

func (ms *MeetingScheduler) CancelMeeting(meetingId int) error {
	if meetingId > len(ms.meeting) {
		return errors.New("There is no such meeting")
	}

	// cancel meeting
	ms.meeting[meetingId].CancelMeeting()

	return nil
}

// GetFreeRoom returns all rooms which are free in particular time interval
func (ms *MeetingScheduler) GetFreeRoom(dur *interval) []*MeetingRoom {
	freeRooms := []*MeetingRoom{}
	for _, room := range ms.rooms {
		if room.IsFree(dur) {
			freeRooms = append(freeRooms, room)
		}

	}

	return freeRooms
}
