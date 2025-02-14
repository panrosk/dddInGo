package storage

import (
	"coworking/internal/app/domain/entities"
	"coworking/internal/app/domain/vo"
	"coworking/internal/ports"
)

type HotDeskRepository struct {
	hotdesk []*entities.Hotdesk
}

func NewHotDeskRepository() *HotDeskRepository {
	return &HotDeskRepository{
		hotdesk: make([]*entities.Hotdesk, 0),
	}
}

func (r *HotDeskRepository) Save(hd *entities.Hotdesk) error {
	if hd == nil {
		return nil
	}
	copy := *hd
	r.hotdesk = append(r.hotdesk, &copy)
	return nil
}

func (r *HotDeskRepository) FindAll() ([]*entities.Hotdesk, error) {
	return r.hotdesk, nil
}

func filter(hotdesks []*entities.Hotdesk, predicate func(*entities.Hotdesk) bool) []*entities.Hotdesk {
	var result []*entities.Hotdesk
	for _, hd := range hotdesks {
		if predicate(hd) {
			result = append(result, hd)
		}
	}
	return result
}

func (r *HotDeskRepository) FindById(number *vo.HotdeskNumber) (*entities.Hotdesk, error) {
	if number == nil {
		return nil, nil
	}
	result := filter(r.hotdesk, func(hd *entities.Hotdesk) bool {
		return hd.Number.Value() == number.Value()
	})
	if len(result) > 0 {
		return result[0], nil
	}
	return nil, nil
}

var _ ports.RepositoryHotdeskPort = (*HotDeskRepository)(nil)
