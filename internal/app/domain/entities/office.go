package entities

import (
	"coworking/internal/app/domain/vo"
	"github.com/google/uuid"
	"time"
)

type OfficeStatus string

type Office struct {
	id          uuid.UUID
	number      vo.OfficeNumber
	status      vo.Status
	leasePeriod vo.OfficeLeasePeriod
	createdAt   time.Time
	updatedAt   time.Time
}

func NewOffice(number int, leasePeriod int, status string) (*Office, error) {
	officeNumber, err := vo.NewOfficeNumber(number)
	if err != nil {
		return nil, err
	}

	officeLeasePeriod, err := vo.NewOfficeLeasePeriod(leasePeriod)
	if err != nil {
		return nil, err
	}

	if status == "" {
		status = "Active"
	}

	parsedStatus, err := vo.NetStatus(status)

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

func (o *Office) GetOffice() map[string]interface{} {
	return map[string]interface{}{
		"id":           o.id.String(),
		"number":       o.number.Value(),
		"status":       string(o.status),
		"lease_period": o.leasePeriod.Value(),
		"created_at":   o.createdAt.Format(time.RFC3339),
		"updated_at":   o.updatedAt.Format(time.RFC3339),
	}
}
