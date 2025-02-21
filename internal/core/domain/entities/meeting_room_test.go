package entities_test

import (
	"coworking/internal/app/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMeetingRoom_Success(t *testing.T) {
	meetingRoom, err := entities.NewMeetingRoom("Conference Room", 10)

	assert.NoError(t, err)
	assert.NotNil(t, meetingRoom)
	assert.Equal(t, "Conference Room", meetingRoom.GetMeetingRoom()["name"])
	assert.Equal(t, 10, meetingRoom.GetMeetingRoom()["capacity"])
	assert.Equal(t, "Active", meetingRoom.GetMeetingRoom()["status"])
	assert.NotEmpty(t, meetingRoom.GetMeetingRoom()["id"])
	assert.NotEmpty(t, meetingRoom.GetMeetingRoom()["created_at"])
	assert.NotEmpty(t, meetingRoom.GetMeetingRoom()["updated_at"])
}

func TestNewMeetingRoom_InvalidName(t *testing.T) {
	meetingRoom, err := entities.NewMeetingRoom("", 10)

	assert.Error(t, err)
	assert.Nil(t, meetingRoom)
}

func TestNewMeetingRoom_InvalidCapacity(t *testing.T) {
	meetingRoom, err := entities.NewMeetingRoom("Small Room", -1)

	assert.Error(t, err)
	assert.Nil(t, meetingRoom)
}
