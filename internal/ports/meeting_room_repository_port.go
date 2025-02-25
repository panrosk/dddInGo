package ports

import "coworking/internal/spaces/meeting_room"

type MeetingRoomRepositoryPort interface {
	RepositoryPort[*meetingroom.MeetingRoom]
	FindByName(room *meetingroom.MeetingRoom) (*meetingroom.MeetingRoom, error)
}
