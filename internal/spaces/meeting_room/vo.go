package meetingroom

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	if value == "" {
		return Name{}, ErrInvalidMeetingRoomName
	}
	return Name{value: value}, nil
}

func (m Name) Value() string {
	return m.value
}

type Capacity struct {
	value int
}

func NewCapacity(value int) (Capacity, error) {
	if value <= 0 {
		return Capacity{}, ErrInvalidMeetingRoomCapacity
	}
	return Capacity{value: value}, nil
}

func (m Capacity) Value() int {
	return m.value
}
