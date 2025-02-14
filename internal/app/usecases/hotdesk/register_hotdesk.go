package hotdesk

import (
	"coworking/internal/app/domain/entities"
	"coworking/internal/app/usecases"
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
	hotdesk, err := entities.NewHotdesk(params.number)
	if err != nil {
		return nil, err
	}

	err = u.storage.SaveHotdesk(hotdesk)

	if err != nil {
		return nil, err
	}

	return hotdesk, nil
}

var _ usecases.Command[RegisterHotdeskParams, *entities.Hotdesk] = (*RegisterHotdeskUsecase)(nil)
