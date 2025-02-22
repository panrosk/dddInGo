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
	storage ports.RepositoryPort[*meetingroom.MeetingRoom]
}

func NewRegisterMeetingRoomUsecase(storage ports.RepositoryPort[*meetingroom.MeetingRoom]) *RegisterMeetingRoomUsecase {
	return &RegisterMeetingRoomUsecase{storage: storage}
}

func (u *RegisterMeetingRoomUsecase) Handle(params RegisterMeetingRoomParams) error {
	newMeetingRoom, err := meetingroom.NewMeetingRoom(params.Name, params.Capacity)
	if err != nil {
		return err
	}

	existingMeetingRooms, err := u.storage.FindByFilter(func(mr *meetingroom.MeetingRoom) bool {
		return mr.GetMeetingRoom()["name"] == params.Name
	})
	if err != nil {
		return err
	}

	if len(existingMeetingRooms) > 0 {
		return meetingroom.ErrMeetingRoomAlreadyExists
	}

	return u.storage.Save(newMeetingRoom)
}
