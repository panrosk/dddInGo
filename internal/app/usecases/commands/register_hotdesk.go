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
	storage ports.RepositoryHotdeskPort
}

func NewRegisterHotdeskUsecase(storage ports.RepositoryHotdeskPort) *RegisterHotdeskUsecase {
	return &RegisterHotdeskUsecase{storage: storage}
}

func (u *RegisterHotdeskUsecase) Execute(params RegisterHotdeskParams) (*entities.Hotdesk, error) {
	hotdesk, err := entities.NewHotdesk(params.Number)
	if err != nil {
		return nil, err
	}

	currentHotdesk, err := u.storage.FindById(&hotdesk.Number)

	if currentHotdesk != nil {
		return nil, domain_errors.ErrHotDeskAlreadyExists
	}

	err = u.storage.Save(hotdesk)

	if err != nil {
		return nil, err
	}

	return hotdesk, nil
}
