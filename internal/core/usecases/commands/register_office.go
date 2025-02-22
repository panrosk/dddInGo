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
	storage ports.RepositoryPort[*office.Office]
}

func NewRegisterOfficeUsecase(storage ports.RepositoryPort[*office.Office]) *RegisterOfficeUsecase {
	return &RegisterOfficeUsecase{storage: storage}
}

func (u *RegisterOfficeUsecase) Handle(params RegisterOfficeParams) error {
	existingOffices, err := u.storage.FindByFilter(func(o *office.Office) bool {
		return o.GetOffice()["number"] == params.Number
	})

	if err != nil {
		return err
	}

	if len(existingOffices) > 0 {
		return office.ErrOfficeAlreadyExists
	}

	newOffice, err := office.NewOffice(params.Number, params.LeasePeriod, params.Status)
	if err != nil {
		return err
	}

	if err := u.storage.Save(newOffice); err != nil {
		return err
	}

	return nil
}
