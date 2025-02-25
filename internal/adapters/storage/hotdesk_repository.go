package storage

import (
	"coworking/internal/ports"
	"coworking/internal/spaces/hotdesk"
	"errors"
)

type HotDeskRepository struct {
	hotdesks []*hotdesk.Hotdesk
}

var _ ports.HotDeskRepositoryPort = (*HotDeskRepository)(nil)

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

func (r *HotDeskRepository) FindHotdeskByNumber(hd *hotdesk.Hotdesk) (*hotdesk.Hotdesk, error) {
	if hd == nil {
		return nil, errors.New("hotdesk cannot be nil")
	}

	for _, storedHd := range r.hotdesks {
		if storedHd.ToMap()["number"] == hd.ToMap()["number"] {
			return storedHd, nil
		}
	}
	return nil, nil
}
