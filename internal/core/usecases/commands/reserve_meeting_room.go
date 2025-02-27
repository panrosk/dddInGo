package commands

import (
	"coworking/internal/ports"
	meetingroom "coworking/internal/spaces/meeting_room"

	"github.com/google/uuid"
)

type ReserveMeetingRoomParams struct {
	MeetingRoom string
	Date        string
}

type ReserveMeetingRoomUseCase struct {
	reservationstorage ports.MeetingRoomReservationRepositoryPort
	meetingroomstorage ports.MeetingRoomRepositoryPort
}

func (u *ReserveMeetingRoomUseCase) Handle(params ReserveMeetingRoomParams) error {
	meetingRoomId, err := uuid.Parse(params.MeetingRoom)
	if err != nil {
		return meetingroom.ErrInvdalidMeetingRoomUUID
	}

	if !u.meetingRoomExists(meetingRoomId) {
		return meetingroom.ErrMeetingRoomNotFound
	}

	date, err := meetingroom.NewDate(params.Date)

	if u.reservationAlreadyExists(meetingRoomId, date) {
		return meetingroom.ErrMeetingRoomAlreadyExists
	}

	meetingRoomReservation, err := meetingroom.NewReservation(meetingRoomId, date)
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
