package commands

import (
	"coworking/internal/ports"
	"coworking/internal/spaces/hotdesk"
)

type RegisterHotdeskParams struct {
	Number int
}

type RegisterHotdeskUsecase struct {
	storage ports.HotDeskRepositoryPort
}

func NewRegisterHotdeskUsecase(storage ports.HotDeskRepositoryPort) *RegisterHotdeskUsecase {
	return &RegisterHotdeskUsecase{storage: storage}
}

func (u *RegisterHotdeskUsecase) Handle(params RegisterHotdeskParams) error {
	newHotdesk, err := createHotdesk(params.Number)
	if err != nil {
		return err
	}

	if u.hotdeskAlreadyExists(newHotdesk) {
		return hotdesk.ErrHotDeskAlreadyExists
	}

	return u.storage.Save(newHotdesk)
}

func createHotdesk(number int) (*hotdesk.Hotdesk, error) {
	return hotdesk.New(number)
}

func (u *RegisterHotdeskUsecase) hotdeskAlreadyExists(hd *hotdesk.Hotdesk) bool {
	existingHotdesk, err := u.storage.FindHotdeskByNumber(hd)
	return err == nil && existingHotdesk != nil
}
