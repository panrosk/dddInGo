package commands_test

import (
	"coworking/internal/core/usecases/commands"
	"coworking/internal/ports"
	"coworking/internal/spaces/hotdesk"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHotdeskStorage struct {
	hotdesks []*hotdesk.Hotdesk
}

func NewMockHotdeskStorage() *MockHotdeskStorage {
	return &MockHotdeskStorage{hotdesks: make([]*hotdesk.Hotdesk, 0)}
}

func (m *MockHotdeskStorage) Save(h *hotdesk.Hotdesk) error {
	if h == nil {
		return errors.New("cannot save nil hotdesk")
	}
	m.hotdesks = append(m.hotdesks, h)
	return nil
}

func (m *MockHotdeskStorage) FindById(id any) (*hotdesk.Hotdesk, error) {
	for _, h := range m.hotdesks {
		hotdeskData := h.GetHotdesk()
		if hotdeskData["id"] == id {
			return h, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *MockHotdeskStorage) FindAll() ([]*hotdesk.Hotdesk, error) {
	return m.hotdesks, nil
}

func (m *MockHotdeskStorage) FindByFilter(filterFunc func(*hotdesk.Hotdesk) bool) ([]*hotdesk.Hotdesk, error) {
	if filterFunc == nil {
		return nil, errors.New("filter function cannot be nil")
	}

	var result []*hotdesk.Hotdesk
	for _, h := range m.hotdesks {
		if filterFunc(h) {
			result = append(result, h)
		}
	}
	return result, nil
}

var _ ports.RepositoryPort[*hotdesk.Hotdesk] = (*MockHotdeskStorage)(nil)

func TestRegisterHotdesk_Success(t *testing.T) {
	mockStorage := NewMockHotdeskStorage()
	usecase := commands.NewRegisterHotdeskUsecase(mockStorage)

	params := commands.RegisterHotdeskParams{
		Number: 5,
	}

	err := usecase.Handle(params)

	assert.NoError(t, err)

	savedHotdesks, _ := mockStorage.FindByFilter(func(h *hotdesk.Hotdesk) bool {
		return h.Number.Value() == params.Number
	})

	assert.Len(t, savedHotdesks, 1)
	assert.Equal(t, params.Number, savedHotdesks[0].Number.Value())
}

func TestRegisterHotdesk_Duplicate(t *testing.T) {
	mockStorage := NewMockHotdeskStorage()
	usecase := commands.NewRegisterHotdeskUsecase(mockStorage)

	existingHotdesk, _ := hotdesk.NewHotdesk(5)
	_ = mockStorage.Save(existingHotdesk)

	params := commands.RegisterHotdeskParams{
		Number: 5,
	}

	err := usecase.Handle(params)

	assert.Error(t, err)
	assert.Equal(t, hotdesk.ErrHotDeskAlreadyExists, err)

	savedHotdesks, _ := mockStorage.FindByFilter(func(h *hotdesk.Hotdesk) bool {
		return h.Number.Value() == params.Number
	})

	assert.Len(t, savedHotdesks, 1)
}

func TestRegisterHotdesk_InvalidNumber(t *testing.T) {
	mockStorage := NewMockHotdeskStorage()
	usecase := commands.NewRegisterHotdeskUsecase(mockStorage)

	params := commands.RegisterHotdeskParams{
		Number: 0,
	}

	err := usecase.Handle(params)

	assert.Error(t, err)

	savedHotdesks, _ := mockStorage.FindAll()
	assert.Empty(t, savedHotdesks)
}
