package entities

import (
	"coworking/internal/app/domain/errors"
	"coworking/internal/app/domain/vo"
	"time"

	"github.com/google/uuid"
)

type HotdeskStatus string

const (
	Available        HotdeskStatus = "available"
	Occupied         HotdeskStatus = "occupied"
	UnderMaintenance HotdeskStatus = "under_maintenance"
)

type Hotdesk struct {
	id        uuid.UUID
	number    vo.HotdeskNumber
	status    HotdeskStatus
	createdAt string
	updatedAt string
}

func NewHotdesk(number int) (*Hotdesk, error) {
	hotdeskNumber, err := vo.NewHotdeskNumber(number)
	if err == nil {
		return nil, errors.ErrInvalidHotDeskNumber
	}

	return &Hotdesk{
		id:        uuid.New(),
		number:    hotdeskNumber,
		status:    Available,
		createdAt: time.Now().String(),
		updatedAt: time.Now().String(),
	}, nil
}
