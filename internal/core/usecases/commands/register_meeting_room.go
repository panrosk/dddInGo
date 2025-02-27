package commands

import (
	"coworking/internal/ports"
	"coworking/internal/spaces/meeting_room"
)

type RegisterMeetingRoomParams struct {
	Name     string
	Capacity int
}

type RegisterMeetingRoomUsecase struct {
	storage ports.MeetingRoomRepositoryPort
}

func NewRegisterMeetingRoomUsecase(storage ports.MeetingRoomRepositoryPort) *RegisterMeetingRoomUsecase {
	return &RegisterMeetingRoomUsecase{storage: storage}
}

func (u *RegisterMeetingRoomUsecase) Handle(params RegisterMeetingRoomParams) error {
	name, err := meetingroom.NewName(params.Name)
	if err != nil {
		return err
	}

	if u.roomAlreadyExists(&name) {
		return meetingroom.ErrMeetingRoomAlreadyExists
	}

	newMeetingRoom, err := createMeetingRoom(params.Name, params.Capacity)
	if err != nil {
		return err
	}

	return u.storage.Save(newMeetingRoom)
}

func createMeetingRoom(name string, capacity int) (*meetingroom.MeetingRoom, error) {
	return meetingroom.New(name, capacity)
}

func (u *RegisterMeetingRoomUsecase) roomAlreadyExists(room *meetingroom.Name) bool {
	existingRoom, err := u.storage.FindByName(room)
	return err == nil && existingRoom != nil
}
