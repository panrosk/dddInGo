package commands

import (
	"coworking/internal/app/domain/domain_errors"
	"coworking/internal/app/domain/entities"
	"coworking/internal/ports"
)

type RegisterMeetingRoomParams struct {
	Name     string
	Capacity int
}

type RegisterMeetingRoomUsecase struct {
	storage ports.RepositoryPort[*entities.MeetingRoom]
}

func NewRegisterMeetingRoomUsecase(storage ports.RepositoryPort[*entities.MeetingRoom]) *RegisterMeetingRoomUsecase {
	return &RegisterMeetingRoomUsecase{storage: storage}
}

func (u *RegisterMeetingRoomUsecase) Execute(params RegisterMeetingRoomParams) (*entities.MeetingRoom, error) {

	meetingRoom, err := entities.NewMeetingRoom(params.Name, params.Capacity)

	existingMeetingRooms, err := u.storage.FindByFilter(func(mr *entities.MeetingRoom) bool {
		return mr.GetMeetingRoom()["name"] == params.Name
	})

	if err != nil {
		return nil, err
	}

	if len(existingMeetingRooms) > 0 {
		return nil, domain_errors.ErrMeetingRoomAlreadyExists
	}

	err = u.storage.Save(meetingRoom)
	if err != nil {
		return nil, err
	}

	return meetingRoom, nil
}
