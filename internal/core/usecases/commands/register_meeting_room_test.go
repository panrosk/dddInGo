package commands_test

import (
	"coworking/internal/app/domain/domain_errors"
	"coworking/internal/app/domain/entities"
	"coworking/internal/app/usecases/commands"
	"coworking/internal/ports"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockMeetingRoomStorage struct {
	meetingRooms []*entities.MeetingRoom
}

func NewMockMeetingRoomStorage() *MockMeetingRoomStorage {
	return &MockMeetingRoomStorage{meetingRooms: make([]*entities.MeetingRoom, 0)}
}

func (m *MockMeetingRoomStorage) Save(mr *entities.MeetingRoom) error {
	m.meetingRooms = append(m.meetingRooms, mr)
	return nil
}

func (m *MockMeetingRoomStorage) FindById(id any) (*entities.MeetingRoom, error) {
	for _, mr := range m.meetingRooms {
		if mr.GetMeetingRoom()["id"] == id {
			return mr, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *MockMeetingRoomStorage) FindAll() ([]*entities.MeetingRoom, error) {
	return m.meetingRooms, nil
}

func (m *MockMeetingRoomStorage) FindByFilter(filterFunc func(*entities.MeetingRoom) bool) ([]*entities.MeetingRoom, error) {
	var result []*entities.MeetingRoom
	for _, mr := range m.meetingRooms {
		if filterFunc(mr) {
			result = append(result, mr)
		}
	}
	return result, nil
}

var _ ports.RepositoryPort[*entities.MeetingRoom] = (*MockMeetingRoomStorage)(nil)

func TestRegisterMeetingRoom_Success(t *testing.T) {
	mockStorage := NewMockMeetingRoomStorage()
	usecase := commands.NewRegisterMeetingRoomUsecase(mockStorage)

	params := commands.RegisterMeetingRoomParams{
		Name:     "Board Room",
		Capacity: 10,
	}

	meetingRoom, err := usecase.Execute(params)

	assert.NoError(t, err)
	assert.NotNil(t, meetingRoom)
	assert.Equal(t, params.Name, meetingRoom.GetMeetingRoom()["name"])
	assert.Equal(t, params.Capacity, meetingRoom.GetMeetingRoom()["capacity"])
}

func TestRegisterMeetingRoom_Duplicate(t *testing.T) {
	mockStorage := NewMockMeetingRoomStorage()
	usecase := commands.NewRegisterMeetingRoomUsecase(mockStorage)

	existingMeetingRoom, _ := entities.NewMeetingRoom("Board Room", 10)
	mockStorage.Save(existingMeetingRoom)

	params := commands.RegisterMeetingRoomParams{
		Name:     "Board Room",
		Capacity: 12,
	}

	meetingRoom, err := usecase.Execute(params)

	assert.Error(t, err)
	assert.Nil(t, meetingRoom)
	assert.Equal(t, domain_errors.ErrMeetingRoomAlreadyExists, err)
}

func TestRegisterMeetingRoom_InvalidName(t *testing.T) {
	mockStorage := NewMockMeetingRoomStorage()
	usecase := commands.NewRegisterMeetingRoomUsecase(mockStorage)

	params := commands.RegisterMeetingRoomParams{
		Name:     "",
		Capacity: 10,
	}

	meetingRoom, err := usecase.Execute(params)

	assert.Error(t, err)
	assert.Nil(t, meetingRoom)
}

func TestRegisterMeetingRoom_InvalidCapacity(t *testing.T) {
	mockStorage := NewMockMeetingRoomStorage()
	usecase := commands.NewRegisterMeetingRoomUsecase(mockStorage)

	params := commands.RegisterMeetingRoomParams{
		Name:     "Conference Room",
		Capacity: -5,
	}

	meetingRoom, err := usecase.Execute(params)

	assert.Error(t, err)
	assert.Nil(t, meetingRoom)
}
