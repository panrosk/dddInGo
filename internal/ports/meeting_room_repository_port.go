package ports

import (
	"coworking/internal/spaces/meeting_room"

	"github.com/google/uuid"
)

type MeetingRoomRepositoryPort interface {
	RepositoryPort[*meetingroom.MeetingRoom]
	FindByName(room *meetingroom.Name) (*meetingroom.MeetingRoom, error)
	FindById(id uuid.UUID) (*meetingroom.MeetingRoom, error)
}
