package ports

import (
	"coworking/internal/spaces/meeting_room"
	"github.com/google/uuid"
)

type MeetingRoomReservationRepositoryPort interface {
	Save(reservation *meetingroom.Reservation) error
	FindByMeetingRoomAndDate(meetingRoomId uuid.UUID, date meetingroom.Date) ([]*meetingroom.Reservation, error)
}
