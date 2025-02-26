package meetingroom

import (
	"coworking/internal/spaces/common"
	"github.com/google/uuid"
	"time"
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

type Reservation struct {
	id            uuid.UUID
	meetingRoomId uuid.UUID
	userId        uuid.UUID
	date          string
	hour          Hour
	duration      Duration
	status        common.Status
	createdAt     time.Time
	updatedAt     time.Time
}

func NewReservation(meetingRoomId, userId uuid.UUID, date string, hour int, duration int) (*Reservation, error) {
	if _, err := time.Parse("2006-01-02", date); err != nil {
		return nil, ErrInvalidDate
	}

	h, err := NewHour(hour)
	if err != nil {
		return nil, err
	}

	d, err := NewDuration(duration)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Reservation{
		id:            uuid.New(),
		meetingRoomId: meetingRoomId,
		userId:        userId,
		date:          date,
		hour:          h,
		duration:      d,
		status:        common.Status("Active"),
		createdAt:     now,
		updatedAt:     now,
	}, nil
}

func (r *Reservation) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            r.id.String(),
		"meetingRoomId": r.meetingRoomId.String(),
		"userId":        r.userId.String(),
		"date":          r.date,
		"hour":          r.hour.Value(),
		"duration":      r.duration.Value(),
		"status":        string(r.status),
		"createdAt":     r.createdAt.Format(time.RFC3339),
		"updatedAt":     r.updatedAt.Format(time.RFC3339),
	}
}
