package storage

import (
	"coworking/internal/ports"
	"coworking/internal/spaces/hotdesk"
	"errors"
)

type HotDeskRepository struct {
	hotdesks []*hotdesk.Hotdesk
}

func NewHotDeskRepository() *HotDeskRepository {
	return &HotDeskRepository{
		hotdesks: make([]*hotdesk.Hotdesk, 0),
	}
}

func (r *HotDeskRepository) Save(hd *hotdesk.Hotdesk) error {
	if hd == nil {
		return errors.New("cannot save a nil hotdesk")
	}
	r.hotdesks = append(r.hotdesks, hd)
	return nil
}

func (r *HotDeskRepository) FindAll() ([]*hotdesk.Hotdesk, error) {
	return r.hotdesks, nil
}

// FindById retrieves a hotdesk by its ID
func (r *HotDeskRepository) FindById(id any) (*hotdesk.Hotdesk, error) {
	number, ok := id.(*hotdesk.Number)
	if !ok {
		return nil, errors.New("invalid ID type, expected *hotdesk.HotdeskNumber")
	}
	if number == nil {
		return nil, errors.New("ID cannot be nil")
	}

	for _, hd := range r.hotdesks {
		if hd.Number.Value() == number.Value() {
			return hd, nil
		}
	}
	return nil, nil
}

func (r *HotDeskRepository) FindByFilter(filterFunc func(*hotdesk.Hotdesk) bool) ([]*hotdesk.Hotdesk, error) {
	if filterFunc == nil {
		return nil, errors.New("filter function cannot be nil")
	}

	var result []*hotdesk.Hotdesk
	for _, hd := range r.hotdesks {
		if filterFunc(hd) {
			result = append(result, hd)
		}
	}
	return result, nil
}

var _ ports.RepositoryPort[*hotdesk.Hotdesk] = (*HotDeskRepository)(nil)
