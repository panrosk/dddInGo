package commands

import (
	"coworking/internal/ports"
	meetingroom "coworking/internal/spaces/meeting_room"
	"time"

	"github.com/google/uuid"
)

type ReserveMeetingRoomParams struct {
	MeetingRoom string
	UserId      string
	Date        string
	Hour        int
	Duration    int
}

type ReserveMeetingRoomUseCase struct {
	reservationstorage        ports.MeetingRoomReservationRepositoryPort
	meetingroomstorage        ports.MeetingRoomRepositoryPort
	reservationhotdeskstorage ports.HotDeskReservationRepositoryPort
}

func (u *ReserveMeetingRoomUseCase) Handle(params ReserveMeetingRoomParams) error {
	meetingRoomId, err := uuid.Parse(params.MeetingRoom)
	if err != nil {
		return meetingroom.ErrInvdalidMeetingRoomUUID
	}

	userId, err := uuid.Parse(params.UserId)
	if err != nil {
		return meetingroom.ErrInvalidUserUUID
	}

	if !u.meetingRoomExists(meetingRoomId) {
		return meetingroom.ErrMeetingRoomNotFound
	}

	date, err := meetingroom.NewDate(params.Date)

	if u.reservationAlreadyExists(meetingRoomId, date) {
		return meetingroom.ErrMeetingRoomAlreadyExists
	}

	if !u.isValidReservationHour(params.Hour, date) {
		return meetingroom.ErrInvalidReservationHour
	}

	meetingRoomReservation, err := meetingroom.NewReservation(meetingRoomId, userId, params.Date, params.Hour, params.Duration)

	if err != nil {
		return err
	}

	return u.reservationstorage.Save(meetingRoomReservation)
}

func (u *ReserveMeetingRoomUseCase) reservationAlreadyExists(id uuid.UUID, date meetingroom.Date) bool {
	existingReservations, err := u.reservationstorage.FindByMeetingRoomAndDate(id, date)

	if err != nil {
		return false
	}

	return len(existingReservations) > 0
}

func (u *ReserveMeetingRoomUseCase) meetingRoomExists(meetingRoomId uuid.UUID) bool {
	_, err := u.meetingroomstorage.FindById(meetingRoomId)
	return err == nil
}

func (u *ReserveMeetingRoomUseCase) isValidReservationHour(hour int, reservationDate meetingroom.Date) bool {
	now := time.Now()

	if now.Format("2006-01-02") == reservationDate.Value() {
		nextValidHour := now.Hour() + 1
		return hour >= nextValidHour
	}

	return false

}
