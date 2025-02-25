package storage

import (
	"coworking/internal/spaces/meeting_room"
	"errors"
)

var ErrMeetingRoomNotFound = errors.New("meeting room not found")

type MeetingRoomRepository struct {
	rooms []*meetingroom.MeetingRoom
}

func NewMeetingRoomRepository() *MeetingRoomRepository {
	return &MeetingRoomRepository{
		rooms: make([]*meetingroom.MeetingRoom, 0),
	}
}

func (r *MeetingRoomRepository) Save(room *meetingroom.MeetingRoom) error {
	if room == nil {
		return errors.New("meeting room cannot be nil")
	}
	r.rooms = append(r.rooms, room)
	return nil
}

func (r *MeetingRoomRepository) FindAll() ([]*meetingroom.MeetingRoom, error) {
	return r.rooms, nil
}

func (r *MeetingRoomRepository) FindByName(room *meetingroom.MeetingRoom) (*meetingroom.MeetingRoom, error) {
	if room == nil {
		return nil, errors.New("meeting room cannot be nil")
	}

	for _, storedRoom := range r.rooms {
		if storedRoom.ToMap()["name"] == room.ToMap()["name"] {
			return storedRoom, nil
		}
	}
	return nil, ErrMeetingRoomNotFound
}
