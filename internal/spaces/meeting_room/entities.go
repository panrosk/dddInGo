package meetingroom

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MeetingRoom struct {
	id        uuid.UUID
	name      Name
	capacity  Capacity
	status    vo.Status
	createdAt time.Time
	updatedAt time.Time
}

func NewMeetingRoom(name string, capacity int) (*MeetingRoom, error) {
	meetingRoomName, err := NewName(name)
	if err != nil {

		fmt.Println("err", err)
		return nil, err
	}

	meetingRoomCapacity, err := NewCapacity(capacity)
	if err != nil {
		return nil, err
	}

	status, err := vo.NewStatus("Active")
	if err != nil {
		return nil, err
	}

	return &MeetingRoom{
		id:        uuid.New(),
		name:      meetingRoomName,
		capacity:  meetingRoomCapacity,
		status:    status,
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
