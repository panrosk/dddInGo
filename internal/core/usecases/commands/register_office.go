package commands

import (
	"coworking/internal/ports"
	"coworking/internal/spaces/office"
)

type RegisterOfficeParams struct {
	Number      int
	LeasePeriod int
	Status      string
}

type RegisterOfficeUsecase struct {
	storage ports.OfficeRepositoryPort
}

func NewRegisterOfficeUsecase(storage ports.OfficeRepositoryPort) *RegisterOfficeUsecase {
	return &RegisterOfficeUsecase{storage: storage}
}

func (u *RegisterOfficeUsecase) Handle(params RegisterOfficeParams) error {
	newOffice, err := createOffice(params.Number, params.LeasePeriod, params.Status)
	if err != nil {
		return err
	}

	if u.officeAlreadyExists(newOffice) {
		return office.ErrOfficeAlreadyExists
	}

	return u.storage.Save(newOffice)
}

func createOffice(number int, leasePeriod int, status string) (*office.Office, error) {
	return office.New(number, leasePeriod, status)
}

func (u *RegisterOfficeUsecase) officeAlreadyExists(o *office.Office) bool {
	existingOffice, err := u.storage.FindByNumber(o)
	return err == nil && existingOffice != nil
}
