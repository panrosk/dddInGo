package ports

import (
	"coworking/internal/app/domain/entities"
	"coworking/internal/app/domain/vo"
)

type RepositoryHotdeskPort interface {
	Save(hotdesk *entities.Hotdesk) error
	FindById(number *vo.HotdeskNumber) (*entities.Hotdesk, error)
	FindAll() ([]*entities.Hotdesk, error)
}
