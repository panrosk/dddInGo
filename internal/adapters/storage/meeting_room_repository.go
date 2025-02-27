package storage

import (
	"coworking/internal/spaces/meeting_room"
	"errors"

	"github.com/google/uuid"
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

func (r *MeetingRoomRepository) FindByName(name *meetingroom.Name) (*meetingroom.MeetingRoom, error) {
	if name == nil {
		return nil, errors.New("name cannot be empty")
	}

	for _, storedRoom := range r.rooms {
		if storedRoom.ToMap()["name"] == name.Value() {
			return storedRoom, nil
		}
	}
	return nil, ErrMeetingRoomNotFound
}

func (r *MeetingRoomRepository) FindById(id uuid.UUID) (*meetingroom.MeetingRoom, error) {
	for _, storedRoom := range r.rooms {
		if storedRoom.ToMap()["id"] == id.String() {
			return storedRoom, nil
		}
	}
	return nil, ErrMeetingRoomNotFound
}
