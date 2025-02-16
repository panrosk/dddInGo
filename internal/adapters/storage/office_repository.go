package storage

import (
	"coworking/internal/app/domain/entities"
	"coworking/internal/ports"
	"errors"
)

type OfficeRepository struct {
	offices []*entities.Office
}

func NewOfficeRepository() *OfficeRepository {
	return &OfficeRepository{
		offices: make([]*entities.Office, 0),
	}
}

func (r *OfficeRepository) Save(office *entities.Office) error {
	if office == nil {
		return errors.New("office cannot be nil")
	}
	copy := *office
	r.offices = append(r.offices, &copy)
	return nil
}

func (r *OfficeRepository) FindAll() ([]*entities.Office, error) {
	return r.offices, nil
}

func filterOffices(offices []*entities.Office, predicate func(*entities.Office) bool) []*entities.Office {
	var result []*entities.Office
	for _, office := range offices {
		if predicate(office) {
			result = append(result, office)
		}
	}
	return result
}

func (r *OfficeRepository) FindById(id any) (*entities.Office, error) {
	officeID, ok := id.(string)
	if !ok {
		return nil, errors.New("invalid ID type, expected string")
	}

	result := filterOffices(r.offices, func(office *entities.Office) bool {
		return office.GetOffice()["id"] == officeID
	})

	if len(result) > 0 {
		return result[0], nil
	}
	return nil, nil
}

func (r *OfficeRepository) FindByFilter(filterFunc func(*entities.Office) bool) ([]*entities.Office, error) {
	if filterFunc == nil {
		return nil, errors.New("filter function cannot be nil")
	}
	return filterOffices(r.offices, filterFunc), nil
}

var _ ports.RepositoryPort[*entities.Office] = (*OfficeRepository)(nil)
