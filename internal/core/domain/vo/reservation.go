package vo

import (
	"coworking/internal/app/domain/domain_errors"
	"time"
)

type ReservationDate struct {
	value time.Time
}

func NewReservationDate(date string) (ReservationDate, error) {
	parsedDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return ReservationDate{}, domain_errors.ErrInvalidReservationDateFormat
	}

	year, month, day := parsedDate.Date()
	onlyDate := time.Date(year, month, day, 0, 0, 0, 0, parsedDate.Location())

	return ReservationDate{value: onlyDate}, nil
}

func (r ReservationDate) Value() time.Time {
	return r.value
}

type ReservationHour struct {
	value int
}

func NewReservationHour(time int) (ReservationHour, error) {
	if time < 0 || time > 23 {
		return ReservationHour{}, domain_errors.ErrInvalidReservationTime
	}

	return ReservationHour{value: time}, nil
}

func (r ReservationHour) Value() int {
	return r.value
}

type ReservationDuration struct {
	value int
}

func NewReservationDuration(duration int) (ReservationDuration, error) {
	if duration < 0 || duration > 12 {
		return ReservationDuration{}, domain_errors.ErrInvalidReservationDuration
	}

	return ReservationDuration{value: duration}, nil
}

func (r ReservationDuration) Value() int {
	return r.value
}
