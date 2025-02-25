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

type Reservation struct {
	id                   uuid.UUID
	userId               uuid.UUID
	date                 time.Time
	status               common.Status
	createdAt            time.Time
	updatedAt            time.Time
	includedInMembership bool
}

func NewReservation(userId uuid.UUID, date time.Time, includedInMembership bool) (*Reservation, error) {
	status, err := defaultStatus()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Reservation{
		id:                   uuid.New(),
		userId:               userId,
		date:                 date,
		status:               status,
		createdAt:            now,
		updatedAt:            now,
		includedInMembership: includedInMembership,
	}, nil
}

func (r *Reservation) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                     r.id.String(),
		"user_id":                r.userId.String(),
		"date":                   r.date.Format(time.RFC3339),
		"status":                 string(r.status),
		"created_at":             r.createdAt.Format(time.RFC3339),
		"updated_at":             r.updatedAt.Format(time.RFC3339),
		"included_in_membership": r.includedInMembership,
	}
}

func defaultStatus() (common.Status, error) {
	return common.NewStatus(defaultStatusValue)
}
