package commands_test

import (
	"coworking/internal/core/usecases/commands"
	"coworking/internal/ports"
	meetingroom "coworking/internal/spaces/meeting_room"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockMeetingRoomStorage struct {
	meetingRooms []*meetingroom.MeetingRoom
}

func NewMockMeetingRoomStorage() *MockMeetingRoomStorage {
	return &MockMeetingRoomStorage{meetingRooms: make([]*meetingroom.MeetingRoom, 0)}
}

func (m *MockMeetingRoomStorage) Save(mr *meetingroom.MeetingRoom) error {
	m.meetingRooms = append(m.meetingRooms, mr)

	return nil
}

func (m *MockMeetingRoomStorage) FindById(id any) (*meetingroom.MeetingRoom, error) {
	for _, mr := range m.meetingRooms {
		if mr.GetMeetingRoom()["id"] == id {
			return mr, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *MockMeetingRoomStorage) FindAll() ([]*meetingroom.MeetingRoom, error) {
	return m.meetingRooms, nil
}

func (m *MockMeetingRoomStorage) FindByFilter(filterFunc func(*meetingroom.MeetingRoom) bool) ([]*meetingroom.MeetingRoom, error) {
	var result []*meetingroom.MeetingRoom
	for _, mr := range m.meetingRooms {
		if filterFunc(mr) {
			result = append(result, mr)
		}
	}
	return result, nil
}

var _ ports.RepositoryPort[*meetingroom.MeetingRoom] = (*MockMeetingRoomStorage)(nil)

func TestRegisterMeetingRoom_Success(t *testing.T) {
	mockStorage := NewMockMeetingRoomStorage()
	usecase := commands.NewRegisterMeetingRoomUsecase(mockStorage)

	params := commands.RegisterMeetingRoomParams{
		Name:     "Board Room",
		Capacity: 10,
	}

	err := usecase.Handle(params)

	assert.NoError(t, err)
}

func TestRegisterMeetingRoom_Duplicate(t *testing.T) {
	mockStorage := NewMockMeetingRoomStorage()
	usecase := commands.NewRegisterMeetingRoomUsecase(mockStorage)

	existingMeetingRoom, _ := meetingroom.NewMeetingRoom("Board Room", 10)
	mockStorage.Save(existingMeetingRoom)

	params := commands.RegisterMeetingRoomParams{
		Name:     "Board Room",
		Capacity: 12,
	}

	err := usecase.Handle(params)

	assert.Error(t, err)
	assert.Equal(t, meetingroom.ErrMeetingRoomAlreadyExists, err)
}

func TestRegisterMeetingRoom_InvalidName(t *testing.T) {
	mockStorage := NewMockMeetingRoomStorage()
	usecase := commands.NewRegisterMeetingRoomUsecase(mockStorage)

	params := commands.RegisterMeetingRoomParams{
		Name:     "",
		Capacity: 10,
	}

	err := usecase.Handle(params)

	assert.Error(t, err)
}

func TestRegisterMeetingRoom_InvalidCapacity(t *testing.T) {
	mockStorage := NewMockMeetingRoomStorage()
	usecase := commands.NewRegisterMeetingRoomUsecase(mockStorage)

	params := commands.RegisterMeetingRoomParams{
		Name:     "Conference Room",
		Capacity: -5,
	}

	err := usecase.Handle(params)

	assert.Error(t, err)
}
