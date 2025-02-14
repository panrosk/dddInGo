package entities

import (
	"coworking/internal/app/domain/vo"
	"time"

	"github.com/google/uuid"
)

type MeetingRoomStatus string

type MeetingRoom struct {
	id        uuid.UUID
	name      vo.MeetingRoomName
	capacity  vo.MeetingRoomCapacity
	status    MeetingRoomStatus
	createdAt time.Time
	updatedAt time.Time
}

func NewMeetingRoom(name string, capacity int) (*MeetingRoom, error) {
	meetingRoomName, err := vo.NewMeetingRoomName(name)
	if err != nil {
		return nil, err
	}

	meetingRoomCapacity, err := vo.NewMeetingRoomCapacity(capacity)
	if err != nil {
		return nil, err
	}

	return &MeetingRoom{
		id:        uuid.New(),
		name:      meetingRoomName,
		capacity:  meetingRoomCapacity,
		status:    MeetingRoomStatus(Available),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (m *MeetingRoom) GetMeetingRoom() map[string]interface{} {
	return map[string]interface{}{
		"id":         m.id.String(),
		"name":       m.name.Value(),
		"capacity":   m.capacity.Value(),
		"status":     string(m.status),
		"created_at": m.createdAt.Format(time.RFC3339),
		"updated_at": m.updatedAt.Format(time.RFC3339),
	}
}
