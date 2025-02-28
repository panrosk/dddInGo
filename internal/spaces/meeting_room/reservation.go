package meetingroom

import (
	"coworking/internal/spaces/common"
	"github.com/google/uuid"
	"time"
)

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
