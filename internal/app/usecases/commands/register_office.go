package commands

import (
	"coworking/internal/app/domain/domain_errors"
	"coworking/internal/app/domain/entities"
	"coworking/internal/ports"
)

type RegisterOfficeParams struct {
	Number      int
	LeasePeriod int
	Status      string
}

type RegisterOfficeUsecase struct {
	storage ports.RepositoryPort[*entities.Office]
}

func NewRegisterOfficeUsecase(storage ports.RepositoryPort[*entities.Office]) *RegisterOfficeUsecase {
	return &RegisterOfficeUsecase{storage: storage}
}

func (u *RegisterOfficeUsecase) Execute(params RegisterOfficeParams) (*entities.Office, error) {
	existingOffices, err := u.storage.FindByFilter(func(o *entities.Office) bool {
		return o.GetOffice()["number"] == params.Number
	})

	if err != nil {
		return nil, err
	}

	if len(existingOffices) > 0 {
		return nil, domain_errors.ErrOfficeAlreadyExists
	}

	office, err := entities.NewOffice(params.Number, params.LeasePeriod, params.Status)
	if err != nil {
		return nil, err
	}

	err = u.storage.Save(office)
	if err != nil {
		return nil, err
	}

	return office, nil
}
