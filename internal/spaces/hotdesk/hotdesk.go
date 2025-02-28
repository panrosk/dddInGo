package hotdesk

import (
	"coworking/internal/spaces/common"
	"time"

	"github.com/google/uuid"
)

const defaultStatusValue = "Active"

type Hotdesk struct {
	id        uuid.UUID
	number    Number
	status    common.Status
	createdAt time.Time
	updatedAt time.Time
}

func New(number int) (*Hotdesk, error) {
	hotdeskNumber, err := NewNumber(number)
	if err != nil {
		return nil, err
	}

	status, err := defaultStatus()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Hotdesk{
		id:        uuid.New(),
		number:    hotdeskNumber,
		status:    status,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func (h *Hotdesk) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         h.id.String(),
		"number":     h.number.Value(),
		"status":     string(h.status),
		"created_at": h.createdAt.Format(time.RFC3339),
		"updated_at": h.updatedAt.Format(time.RFC3339),
	}
}

func defaultStatus() (common.Status, error) {
	return common.NewStatus(defaultStatusValue)
}
