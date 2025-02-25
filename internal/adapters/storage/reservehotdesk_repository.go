package storage

import (
	"coworking/internal/spaces/hotdesk"
	"errors"
)

type HotDeskReservationRepository struct {
	reservations []*hotdesk.Reservation
}

func NewHotDeskReservationRepository() *HotDeskReservationRepository {
	return &HotDeskReservationRepository{
		reservations: make([]*hotdesk.Reservation, 0),
	}
}

func (r *HotDeskReservationRepository) Save(reservation *hotdesk.Reservation) error {
	if reservation == nil {
		return errors.New("reservation cannot be nil")
	}
	copy := *reservation
	r.reservations = append(r.reservations, &copy)
	return nil
}

func (r *HotDeskReservationRepository) FindAll() ([]*hotdesk.Reservation, error) {
	return r.reservations, nil
}

func (r *HotDeskReservationRepository) FindByReservation(reservation *hotdesk.Reservation) ([]*hotdesk.Reservation, error) {
	if reservation == nil {
		return nil, errors.New("reservation cannot be nil")
	}

	var result []*hotdesk.Reservation
	for _, res := range r.reservations {
		if res.ToMap()["user_id"] == reservation.ToMap()["user_id"] && res.ToMap()["date"] == reservation.ToMap()["date"] {
			result = append(result, res)
		}
	}
	return result, nil
}
