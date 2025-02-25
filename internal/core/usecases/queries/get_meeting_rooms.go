package queries

import (
	"coworking/internal/ports"
	meetingroom "coworking/internal/spaces/meeting_room"

	"github.com/google/uuid"
)

type GetMeetingRoomByIdQuery struct {
	Id string
}

type GetMeetingRoomByIdUsecase struct {
	meetingRoomRepo ports.MeetingRoomRepositoryPort
}

func NewGetMeetingRoomByIdUsecase(meetingRoomRepo ports.MeetingRoomRepositoryPort) *GetMeetingRoomByIdUsecase {
	return &GetMeetingRoomByIdUsecase{
		meetingRoomRepo: meetingRoomRepo,
	}
}

func (u *GetMeetingRoomByIdUsecase) Handle(query GetMeetingRoomByIdQuery) (*meetingroom.MeetingRoom, error) {
	meetingRoomId, err := uuid.Parse(query.Id)
	if err != nil {
		return nil, meetingroom.ErrInvalidMeetingRoomID
	}

	room, err := u.meetingRoomRepo.FindById(meetingRoomId)
	if err != nil {
		return nil, err
	}

	return room, nil
}
