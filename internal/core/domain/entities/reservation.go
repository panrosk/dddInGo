package entities

import (
	"coworking/internal/app/domain/vo"
	"github.com/google/uuid"
	"time"
)

type Reservation struct {
	id            uuid.UUID
	meetingRoomID uuid.UUID
	userID        uuid.UUID
	date          vo.ReservationDate
	hour          vo.ReservationHour
	duration      vo.ReservationDuration
	status        vo.Status
	createdAt     time.Time
	updatedAt     time.Time
}

func NewReservation(
	meetingRoomID uuid.UUID,
	userID uuid.UUID,
	date string,
	hour int,
	duration int,
	status string) (*Reservation, error) {

	reservationDate, err := vo.NewReservationDate(date)
	if err != nil {
		return nil, err
	}

	reservationHour, err := vo.NewReservationHour(hour)
	if err != nil {
		return nil, err
	}

	reservationDuration, err := vo.NewReservationDuration(duration)
	if err != nil {
		return nil, err
	}

	reservationStatus, err := vo.NewStatus(status) // Error tipográfico corregido
	if err != nil {
		return nil, err
	}

	return &Reservation{
		id:            uuid.New(),
		meetingRoomID: meetingRoomID,
		userID:        userID,
		date:          reservationDate,
		hour:          reservationHour,
		duration:      reservationDuration,
		status:        reservationStatus,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}, nil
}

func (r *Reservation) ID() uuid.UUID            { return r.id }
func (r *Reservation) MeetingRoomID() uuid.UUID { return r.meetingRoomID }
func (r *Reservation) UserID() uuid.UUID        { return r.userID }
func (r *Reservation) Date() time.Time          { return r.date.Value() }
func (r *Reservation) Hour() int                { return r.hour.Value() }
func (r *Reservation) Duration() int            { return r.duration.Value() }
func (r *Reservation) Status() string           { return string(r.status) }
func (r *Reservation) CreatedAt() time.Time     { return r.createdAt }
func (r *Reservation) UpdatedAt() time.Time     { return r.updatedAt }

func (r *Reservation) UpdateStatus(newStatus string) error {
	status, err := vo.NewStatus(newStatus)
	if err != nil {
		return err
	}
	r.status = status
	r.updatedAt = time.Now()
	return nil
}

func (r *Reservation) ExtendDuration(extraMinutes int) error {
	newDuration := r.duration.Value() + extraMinutes
	updatedDuration, err := vo.NewReservationDuration(newDuration)
	if err != nil {
		return err
	}
	r.duration = updatedDuration
	r.updatedAt = time.Now()
	return nil
}

func (r *Reservation) GetReservation() map[string]interface{} {
	return map[string]interface{}{
		"id":              r.id.String(),
		"meeting_room_id": r.meetingRoomID.String(),
		"user_id":         r.userID.String(),
		"date":            r.date.Value(),
		"hour":            r.hour.Value(),
		"duration":        r.duration.Value(),
		"status":          string(r.status),
		"created_at":      r.createdAt.Format(time.RFC3339),
		"updated_at":      r.updatedAt.Format(time.RFC3339),
	}
}
