package vo

import "coworking/internal/app/domain/domain_errors"

type OfficeNumber struct {
	value int
}

func NewOfficeNumber(value int) (OfficeNumber, error) {
	if value <= 0 {
		return OfficeNumber{}, domain_errors.ErrInvalidOfficeNumber
	}
	return OfficeNumber{value: value}, nil
}

func (o OfficeNumber) Value() int {
	return o.value
}

type OfficeLeasePeriod struct {
	value int
}

func NewOfficeLeasePeriod(value int) (OfficeLeasePeriod, error) {
	if value <= 0 {
		return OfficeLeasePeriod{}, domain_errors.ErrInvalidOfficeLeasePeriod
	}
	return OfficeLeasePeriod{value: value}, nil
}

func (o OfficeLeasePeriod) Value() int {
	return o.value
}

type OfficeStatus struct {
	value string
}
