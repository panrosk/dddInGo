package commands

import (
	"coworking/internal/app/domain/domain_errors"
	"coworking/internal/app/domain/entities"
	"coworking/internal/ports"
)

type RegisterHotdeskParams struct {
	Number int
}

type RegisterHotdeskUsecase struct {
	storage ports.RepositoryPort[*entities.Hotdesk]
}

func NewRegisterHotdeskUsecase(storage ports.RepositoryPort[*entities.Hotdesk]) *RegisterHotdeskUsecase {
	return &RegisterHotdeskUsecase{storage: storage}
}

func (u *RegisterHotdeskUsecase) Execute(params RegisterHotdeskParams) (*entities.Hotdesk, error) {
	hotdesk, err := entities.NewHotdesk(params.Number)
	if err != nil {
		return nil, err
	}

	existingHotdesks, err := u.storage.FindByFilter(func(hd *entities.Hotdesk) bool {
		return hd.Number.Value() == params.Number
	})

	if err != nil {
		return nil, err
	}

	if len(existingHotdesks) > 0 {
		return nil, domain_errors.ErrHotDeskAlreadyExists
	}

	err = u.storage.Save(hotdesk)
	if err != nil {
		return nil, err
	}

	return hotdesk, nil
}
