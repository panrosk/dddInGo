package storage

import (
	"coworking/internal/ports"
	meetingroom "coworking/internal/spaces/meeting_room"
	"errors"
)

type MeetingRoomRepository struct {
	meetingRooms []*meetingroom.MeetingRoom
}

func NewMeetingRoomRepository() *MeetingRoomRepository {
	return &MeetingRoomRepository{
		meetingRooms: make([]*meetingroom.MeetingRoom, 0),
	}
}

func (r *MeetingRoomRepository) Save(mr *meetingroom.MeetingRoom) error {
	if mr == nil {
		return errors.New("meeting room cannot be nil")
	}
	copy := *mr
	r.meetingRooms = append(r.meetingRooms, &copy)
	return nil
}

func (r *MeetingRoomRepository) FindAll() ([]*meetingroom.MeetingRoom, error) {
	return r.meetingRooms, nil
}

func filterMeetingRooms(meetingRooms []*meetingroom.MeetingRoom, predicate func(*meetingroom.MeetingRoom) bool) []*meetingroom.MeetingRoom {
	var result []*meetingroom.MeetingRoom
	for _, mr := range meetingRooms {
		if predicate(mr) {
			result = append(result, mr)
		}
	}
	return result
}

func (r *MeetingRoomRepository) FindById(id any) (*meetingroom.MeetingRoom, error) {
	roomID, ok := id.(string)
	if !ok {
		return nil, errors.New("invalid ID type, expected string")
	}

	result := filterMeetingRooms(r.meetingRooms, func(mr *meetingroom.MeetingRoom) bool {
		return mr.GetMeetingRoom()["id"] == roomID
	})

	if len(result) > 0 {
		return result[0], nil
	}
	return nil, nil
}

func (r *MeetingRoomRepository) FindByFilter(filterFunc func(*meetingroom.MeetingRoom) bool) ([]*meetingroom.MeetingRoom, error) {
	if filterFunc == nil {
		return nil, errors.New("filter function cannot be nil")
	}
	return filterMeetingRooms(r.meetingRooms, filterFunc), nil
}

var _ ports.RepositoryPort[*meetingroom.MeetingRoom] = (*MeetingRoomRepository)(nil)
