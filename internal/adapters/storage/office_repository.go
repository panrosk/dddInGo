package storage

import (
	"coworking/internal/ports"
	"coworking/internal/spaces/office"
	"errors"
)

type OfficeRepository struct {
	offices []*office.Office
}

func NewOfficeRepository() *OfficeRepository {
	return &OfficeRepository{
		offices: make([]*office.Office, 0),
	}
}

func (r *OfficeRepository) Save(o *office.Office) error {
	if o == nil {
		return errors.New("office cannot be nil")
	}
	r.offices = append(r.offices, o)
	return nil
}

func (r *OfficeRepository) FindAll() ([]*office.Office, error) {
	return r.offices, nil
}

func (r *OfficeRepository) FindById(id any) (*office.Office, error) {
	officeID, ok := id.(string)
	if !ok {
		return nil, errors.New("invalid ID type, expected string")
	}

	for _, o := range r.offices {
		if o.GetOffice()["id"] == officeID {
			return o, nil
		}
	}
	return nil, nil
}

func (r *OfficeRepository) FindByFilter(filterFunc func(*office.Office) bool) ([]*office.Office, error) {
	if filterFunc == nil {
		return nil, errors.New("filter function cannot be nil")
	}

	var result []*office.Office
	for _, o := range r.offices {
		if filterFunc(o) {
			result = append(result, o)
		}
	}
	return result, nil
}

// Implementación de la interfaz RepositoryPort
var _ ports.RepositoryPort[*office.Office] = (*OfficeRepository)(nil)
