package ports

import (
	"coworking/internal/spaces/hotdesk"
)

type HotDeskReservationRepositoryPort interface {
	RepositoryPort[*hotdesk.Reservation]
	FindByReservation(reservation *hotdesk.Reservation) ([]*hotdesk.Reservation, error)
}
