package meetingroom

import (
	"coworking/internal/spaces/common"
	"time"

	"github.com/google/uuid"
)

type MeetingRoom struct {
	id        uuid.UUID
	name      Name
	capacity  Capacity
	status    common.Status
	createdAt time.Time
	updatedAt time.Time
}

func New(name string, capacity int) (*MeetingRoom, error) {
	meetingRoomName, err := createMeetingRoomName(name)
	if err != nil {
		return nil, err
	}

	meetingRoomCapacity, err := createMeetingRoomCapacity(capacity)
	if err != nil {
		return nil, err
	}

	status, err := defaultMeetingRoomStatus()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &MeetingRoom{
		id:        uuid.New(),
		name:      meetingRoomName,
		capacity:  meetingRoomCapacity,
		status:    status,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func (m *MeetingRoom) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         m.id.String(),
		"name":       m.name.Value(),
		"capacity":   m.capacity.Value(),
		"status":     string(m.status),
		"created_at": formatTimestamp(m.createdAt),
		"updated_at": formatTimestamp(m.updatedAt),
	}
}

func createMeetingRoomName(name string) (Name, error) {
	return NewName(name)
}

func createMeetingRoomCapacity(capacity int) (Capacity, error) {
	return NewCapacity(capacity)
}

func defaultMeetingRoomStatus() (common.Status, error) {
	return common.NewStatus("Active")
}

func formatTimestamp(t time.Time) string {
	return t.Format(time.RFC3339)
}
