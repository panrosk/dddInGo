package commands

import (
	"coworking/internal/ports"
	"coworking/internal/spaces/hotdesk"
)

type RegisterHotdeskParams struct {
	Number int
}

type RegisterHotdeskUsecase struct {
	storage ports.RepositoryPort[*hotdesk.Hotdesk]
}

func NewRegisterHotdeskUsecase(storage ports.RepositoryPort[*hotdesk.Hotdesk]) *RegisterHotdeskUsecase {
	return &RegisterHotdeskUsecase{storage: storage}
}

func (u *RegisterHotdeskUsecase) Handle(params RegisterHotdeskParams) error {
	newHotdesk, err := hotdesk.NewHotdesk(params.Number)
	if err != nil {
		return err
	}

	existingHotdesks, err := u.storage.FindByFilter(func(hd *hotdesk.Hotdesk) bool {
		return hd.Number.Value() == params.Number
	})

	if err != nil {
		return err
	}

	if len(existingHotdesks) > 0 {
		return hotdesk.ErrHotDeskAlreadyExists
	}

	if err := u.storage.Save(newHotdesk); err != nil {
		return err
	}

	return nil
}
