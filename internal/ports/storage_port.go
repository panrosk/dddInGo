package ports

import "coworking/internal/app/domain/entities"

type RepositoryHotdeskPort interface {
	SaveHotdesk(hotdesk *entities.Hotdesk) error
	FindByIdHotdesk(hotdeskNumber int) (entities.Hotdesk, error)
	FindAllHotdesk() ([]*entities.Hotdesk, error)
}
