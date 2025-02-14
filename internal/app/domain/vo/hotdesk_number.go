package vo

import (
	"coworking/internal/app/domain/errors"
)

type HotdeskNumber struct {
	value int
}

func NewHotdeskNumber(value int) (HotdeskNumber, error) {
	if value <= 0 {
		return HotdeskNumber{}, errors.ErrInvalidHotDeskNumber
	}
	return HotdeskNumber{value: value}, nil
}

func (hn HotdeskNumber) Value() int {
	return hn.value
}
