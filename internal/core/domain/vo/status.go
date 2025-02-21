package vo

import "coworking/internal/app/domain/domain_errors"

type Status string

const (
	Active           Status = "Active"
	Inactive         Status = "Inactive"
	Occupied         Status = "Occupied"
	UnderMaintenance Status = "Under_maintenance"
)

func NewStatus(s string) (Status, error) {
	switch Status(s) {
	case Active, Inactive, Occupied, UnderMaintenance:
		return Status(s), nil
	default:
		return "", domain_errors.ErrInvalidStatus
	}
}
