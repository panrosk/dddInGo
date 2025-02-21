package office

type Number struct {
	value int
}

func NewNumber(value int) (Number, error) {
	if value <= 0 {
		return Number{}, ErrInvalidOfficeNumber
	}
	return Number{value: value}, nil
}

func (o Number) Value() int {
	return o.value
}

type LeasePeriod struct {
	value int
}

func NewLeasePeriod(value int) (LeasePeriod, error) {
	if value <= 0 {
		return LeasePeriod{}, ErrInvalidOfficeLeasePeriod
	}
	return LeasePeriod{value: value}, nil
}

func (o LeasePeriod) Value() int {
	return o.value
}
