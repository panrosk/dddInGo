package storage

import (
	"coworking/internal/spaces/office"
	"errors"
)

var ErrOfficeNotFound = errors.New("office not found")

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

func (r *OfficeRepository) FindByNumber(o *office.Office) (*office.Office, error) {
	if o == nil {
		return nil, errors.New("office cannot be nil")
	}

	for _, storedOffice := range r.offices {
		if storedOffice.ToMap()["number"] == o.ToMap()["number"] {
			return storedOffice, nil
		}
	}
	return nil, ErrOfficeNotFound
}
