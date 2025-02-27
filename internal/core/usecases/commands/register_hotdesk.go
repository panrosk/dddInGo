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

	number, err := hotdesk.NewNumber(params.Number)
	if err != nil {
		return err
	}

	if u.hotdeskAlreadyExists(number) {
		return hotdesk.ErrHotDeskAlreadyExists
	}

	newHotdesk, err := createHotdesk(params.Number)
	if err != nil {
		return err
	}

	return u.storage.Save(newHotdesk)
}

func createHotdesk(number int) (*hotdesk.Hotdesk, error) {
	return hotdesk.New(number)
}

func (u *RegisterHotdeskUsecase) hotdeskAlreadyExists(number hotdesk.Number) bool {
	existingHotdesk, err := u.storage.FindHotdeskByNumber(&number)
	return err == nil && existingHotdesk != nil
}
