package hotdesk

import (
	"coworking/internal/spaces/common"
	"time"

	"github.com/google/uuid"
)

type Hotdesk struct {
	id        uuid.UUID
	Number    Number
	status    common.Status
	createdAt time.Time
	updatedAt time.Time
}

func NewHotdesk(number int) (*Hotdesk, error) {
	hotdeskNumber, err := NewNumber(number)
	if err != nil {
		return nil, err
	}

	status, err := common.NewStatus("Active")
	if err != nil {
		return nil, err
	}

	return &Hotdesk{
		id:        uuid.New(),
		Number:    hotdeskNumber,
		status:    status,
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

