package storage

import (
	"coworking/internal/spaces/meeting_room"
	"errors"

	"github.com/google/uuid"
)

var ErrReservationNotFound = errors.New("reservation not found")

type MeetingRoomReservationRepository struct {
	reservations []*meetingroom.Reservation
}

func NewMeetingRoomReservationRepository() *MeetingRoomReservationRepository {
	return &MeetingRoomReservationRepository{
		reservations: make([]*meetingroom.Reservation, 0),
	}
}

func (r *MeetingRoomReservationRepository) Save(reservation *meetingroom.Reservation) error {
	if reservation == nil {
		return errors.New("reservation cannot be nil")
	}
	r.reservations = append(r.reservations, reservation)
	return nil
}

func (r *MeetingRoomReservationRepository) FindByMeetingRoomAndDate(meetingRoomId uuid.UUID, date string) ([]*meetingroom.Reservation, error) {
	var results []*meetingroom.Reservation
	for _, res := range r.reservations {
		resMap := res.ToMap()
		if resMap["meetingRoomId"] == meetingRoomId.String() && resMap["date"] == date {
			results = append(results, res)
		}
	}
	return results, nil
}

func (r *MeetingRoomReservationRepository) FindByUser(userId uuid.UUID) ([]*meetingroom.Reservation, error) {
	var results []*meetingroom.Reservation
	for _, res := range r.reservations {
		resMap := res.ToMap()
		if resMap["userId"] == userId.String() {
			results = append(results, res)
		}
	}
	return results, nil
}
