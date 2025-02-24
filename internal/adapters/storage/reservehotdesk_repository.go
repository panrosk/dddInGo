package storage

import (
	"coworking/internal/ports"
	"coworking/internal/spaces/hotdesk"
	"errors"
)

type HotDeskReservationRepository struct {
	reservations []*hotdesk.HotDeskReservation
}

func NewHotDeskReservationRepository() *HotDeskReservationRepository {
	return &HotDeskReservationRepository{
		reservations: make([]*hotdesk.HotDeskReservation, 0),
	}
}

func (r *HotDeskReservationRepository) Save(reservation *hotdesk.HotDeskReservation) error {
	if reservation == nil {
		return errors.New("reservation cannot be nil")
	}
	copy := *reservation
	r.reservations = append(r.reservations, &copy)
	return nil
}

func (r *HotDeskReservationRepository) FindAll() ([]*hotdesk.HotDeskReservation, error) {

	return r.reservations, nil
}

func filterReservations(reservations []*hotdesk.HotDeskReservation, predicate func(*hotdesk.HotDeskReservation) bool) []*hotdesk.HotDeskReservation {

	var result []*hotdesk.HotDeskReservation
	for _, res := range reservations {
		if predicate(res) {
			result = append(result, res)
		}
	}
	return result
}

func (r *HotDeskReservationRepository) FindById(id any) (*hotdesk.HotDeskReservation, error) {
	reservationID, ok := id.(string)
	if !ok {
		return nil, errors.New("invalid ID type, expected string")
	}

	result := filterReservations(r.reservations, func(res *hotdesk.HotDeskReservation) bool {
		return res.GetReservation()["id"] == reservationID
	})

	if len(result) > 0 {
		return result[0], nil
	}
	return nil, nil
}

func (r *HotDeskReservationRepository) FindByFilter(filterFunc func(*hotdesk.HotDeskReservation) bool) ([]*hotdesk.HotDeskReservation, error) {
	if filterFunc == nil {
		return nil, errors.New("filter function cannot be nil")
	}
	return filterReservations(r.reservations, filterFunc), nil
}

var _ ports.RepositoryPort[*hotdesk.HotDeskReservation] = (*HotDeskReservationRepository)(nil)
