package ports

import (
	"coworking/internal/spaces/hotdesk"
	"github.com/google/uuid"
	"time"
)

type HotDeskReservationRepositoryPort interface {
	RepositoryPort[*hotdesk.Reservation]
	FindByUserIDAndDate(userID uuid.UUID, date time.Time) ([]*hotdesk.Reservation, error)
}
