package meetingroom

import (
	"errors"
	"time"
)

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

type Hour struct {
	value int
}

func NewHour(hour int) (Hour, error) {
	if hour < 0 || hour > 23 {
		return Hour{}, ErrInvalidHours
	}
	return Hour{value: hour}, nil
}

func (h Hour) Value() int {
	return h.value
}

type Duration struct {
	value int
}

func NewDuration(duration int) (Duration, error) {
	if duration < 1 || duration > 12 {
		return Duration{}, ErrInvalidDuration
	}
	return Duration{value: duration}, nil
}

func (d Duration) Value() int {
	return d.value
}

type Date struct {
	value time.Time
}

func NewDate(dateStr string) (Date, error) {
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return Date{}, errors.New("invalid date format, expected YYYY-MM-DD")
	}
	return Date{value: parsedDate}, nil
}

func (d Date) Value() string {
	return d.value.Format("2006-01-02")
}
