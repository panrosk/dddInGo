package hotdesk

type Number struct {
	value int
}

func NewNumber(value int) (Number, error) {
	if !isValidHotDeskNumber(value) {
		return Number{}, ErrInvalidHotDeskNumber
	}
	return Number{value: value}, nil
}

func isValidHotDeskNumber(value int) bool {
	return value > 0
}

func (o Number) Value() int {
	return o.value
}
