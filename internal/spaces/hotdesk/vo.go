package hotdesk

type Number struct {
	value int
}

func NewNumber(value int) (Number, error) {
	if value <= 0 {
		return Number{}, ErrInvalidHotDeskNumber
	}
	return Number{value: value}, nil
}

func (o Number) Value() int {
	return o.value
}
