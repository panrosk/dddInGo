package hotdesk

import (
	"coworking/internal/spaces/common"
	"github.com/google/uuid"
	"time"
)

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
