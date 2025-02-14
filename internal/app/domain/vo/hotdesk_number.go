package vo

import "coworking/internal/app/domain/domain_errors"

type HotdeskNumber struct {
	value int
}

func NewHotdeskNumber(value int) (HotdeskNumber, error) {
	if value < 1 {
		return HotdeskNumber{}, domain_errors.ErrInvalidHotDeskNumber
	}
	return HotdeskNumber{value: value}, nil
}

func (h HotdeskNumber) Value() int {
	return h.value
}
