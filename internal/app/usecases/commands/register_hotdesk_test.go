package commands_test

import (
	"coworking/internal/app/domain/domain_errors"
	"coworking/internal/app/domain/entities"
	"coworking/internal/app/usecases/commands"
	"coworking/internal/ports"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHotdeskStorage struct {
	hotdesks []*entities.Hotdesk
}

func NewMockHotdeskStorage() *MockHotdeskStorage {
	return &MockHotdeskStorage{hotdesks: make([]*entities.Hotdesk, 0)}
}

func (m *MockHotdeskStorage) Save(h *entities.Hotdesk) error {
	m.hotdesks = append(m.hotdesks, h)
	return nil
}

func (m *MockHotdeskStorage) FindById(id any) (*entities.Hotdesk, error) {
	for _, h := range m.hotdesks {
		if h.GetHotdesk()["id"] == id {
			return h, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *MockHotdeskStorage) FindAll() ([]*entities.Hotdesk, error) {
	return m.hotdesks, nil
}

func (m *MockHotdeskStorage) FindByFilter(filterFunc func(*entities.Hotdesk) bool) ([]*entities.Hotdesk, error) {
	var result []*entities.Hotdesk
	for _, h := range m.hotdesks {
		if filterFunc(h) {
			result = append(result, h)
		}
	}
	return result, nil
}

var _ ports.RepositoryPort[*entities.Hotdesk] = (*MockHotdeskStorage)(nil)

func TestRegisterHotdesk_Success(t *testing.T) {
	mockStorage := NewMockHotdeskStorage()
	usecase := commands.NewRegisterHotdeskUsecase(mockStorage)

	params := commands.RegisterHotdeskParams{
		Number: 5,
	}

	hotdesk, err := usecase.Execute(params)

	assert.NoError(t, err)
	assert.NotNil(t, hotdesk)
	assert.Equal(t, params.Number, hotdesk.GetHotdesk()["number"])
}

func TestRegisterHotdesk_Duplicate(t *testing.T) {
	mockStorage := NewMockHotdeskStorage()
	usecase := commands.NewRegisterHotdeskUsecase(mockStorage)

	existingHotdesk, _ := entities.NewHotdesk(5)
	mockStorage.Save(existingHotdesk)

	params := commands.RegisterHotdeskParams{
		Number: 5,
	}

	hotdesk, err := usecase.Execute(params)

	assert.Error(t, err)
	assert.Nil(t, hotdesk)
	assert.Equal(t, domain_errors.ErrHotDeskAlreadyExists, err)
}

func TestRegisterHotdesk_InvalidNumber(t *testing.T) {
	mockStorage := NewMockHotdeskStorage()
	usecase := commands.NewRegisterHotdeskUsecase(mockStorage)

	params := commands.RegisterHotdeskParams{
		Number: 0,
	}

	hotdesk, err := usecase.Execute(params)

	assert.Error(t, err)
	assert.Nil(t, hotdesk)
}
