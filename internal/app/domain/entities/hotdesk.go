package entities

import (
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
	Number    vo.HotdeskNumber
	status    HotdeskStatus
	createdAt time.Time
	updatedAt time.Time
}

func NewHotdesk(number int) (*Hotdesk, error) {
	hotdeskNumber, err := vo.NewHotdeskNumber(number)

	if err != nil {
		return nil, err
	}

	return &Hotdesk{
		id:        uuid.New(),
		Number:    hotdeskNumber,
		status:    Available,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (h *Hotdesk) GetHotdesk() map[string]interface{} {
	return map[string]interface{}{
		"id":         h.id.String(),
		"number":     h.Number.Value(),
		"status":     string(h.status),
		"created_at": h.createdAt.Format(time.RFC3339),
		"updated_at": h.updatedAt.Format(time.RFC3339),
	}
}
