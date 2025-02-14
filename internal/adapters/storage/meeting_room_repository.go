package storage

import (
	"coworking/internal/app/domain/entities"
	"coworking/internal/ports"
	"errors"
)

type MeetingRoomRepository struct {
	meetingRooms []*entities.MeetingRoom
}

func NewMeetingRoomRepository() *MeetingRoomRepository {
	return &MeetingRoomRepository{
		meetingRooms: make([]*entities.MeetingRoom, 0),
	}
}

func (r *MeetingRoomRepository) Save(mr *entities.MeetingRoom) error {
	if mr == nil {
		return errors.New("meeting room cannot be nil")
	}
	copy := *mr
	r.meetingRooms = append(r.meetingRooms, &copy)
	return nil
}

func (r *MeetingRoomRepository) FindAll() ([]*entities.MeetingRoom, error) {
	return r.meetingRooms, nil
}

func filterMeetingRooms(meetingRooms []*entities.MeetingRoom, predicate func(*entities.MeetingRoom) bool) []*entities.MeetingRoom {
	var result []*entities.MeetingRoom
	for _, mr := range meetingRooms {
		if predicate(mr) {
			result = append(result, mr)
		}
	}
	return result
}

func (r *MeetingRoomRepository) FindById(id any) (*entities.MeetingRoom, error) {
	roomID, ok := id.(string)
	if !ok {
		return nil, errors.New("invalid ID type, expected string")
	}

	result := filterMeetingRooms(r.meetingRooms, func(mr *entities.MeetingRoom) bool {
		return mr.GetMeetingRoom()["id"] == roomID
	})

	if len(result) > 0 {
		return result[0], nil
	}
	return nil, nil
}

func (r *MeetingRoomRepository) FindByFilter(filterFunc func(*entities.MeetingRoom) bool) ([]*entities.MeetingRoom, error) {
	if filterFunc == nil {
		return nil, errors.New("filter function cannot be nil")
	}
	return filterMeetingRooms(r.meetingRooms, filterFunc), nil
}

var _ ports.RepositoryPort[*entities.MeetingRoom] = (*MeetingRoomRepository)(nil)
