package office

import (
	"coworking/internal/spaces/common"
	"time"

	"github.com/google/uuid"
)

type Office struct {
	id          uuid.UUID
	number      Number
	status      common.Status
	leasePeriod LeasePeriod
	createdAt   time.Time
	updatedAt   time.Time
}

func New(number int, leasePeriod int, status string) (*Office, error) {
	officeNumber, err := NewNumber(number)
	if err != nil {
		return nil, err
	}

	officeLeasePeriod, err := NewLeasePeriod(leasePeriod)
	if err != nil {
		return nil, err
	}

	if status == "" {
		status = "Active"
	}

	parsedStatus, err := common.NewStatus(status)
	if err != nil {
		return nil, err
	}

	return &Office{
		id:          uuid.New(),
		number:      officeNumber,
		status:      parsedStatus,
		leasePeriod: officeLeasePeriod,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}, nil
}

func (o *Office) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":           o.id.String(),
		"number":       o.number.Value(),
		"status":       string(o.status),
		"lease_period": o.leasePeriod.Value(),
		"created_at":   o.createdAt.Format(time.RFC3339),
		"updated_at":   o.updatedAt.Format(time.RFC3339),
	}
}
