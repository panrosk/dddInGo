package commands_test

import (
	"coworking/internal/core/usecases/commands"
	"coworking/internal/ports"
	"coworking/internal/spaces/office"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockOfficeStorage struct {
	offices []*office.Office
}

func NewMockOfficeStorage() *MockOfficeStorage {
	return &MockOfficeStorage{offices: make([]*office.Office, 0)}
}

func (m *MockOfficeStorage) Save(o *office.Office) error {
	m.offices = append(m.offices, o)
	return nil
}

func (m *MockOfficeStorage) FindById(id any) (*office.Office, error) {
	for _, o := range m.offices {
		if o.GetOffice()["id"] == id {
			return o, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *MockOfficeStorage) FindAll() ([]*office.Office, error) {
	return m.offices, nil
}

func (m *MockOfficeStorage) FindByFilter(filterFunc func(*office.Office) bool) ([]*office.Office, error) {
	var result []*office.Office
	for _, o := range m.offices {
		if filterFunc(o) {
			result = append(result, o)
		}
	}
	return result, nil
}

var _ ports.RepositoryPort[*office.Office] = (*MockOfficeStorage)(nil)

func TestRegisterOffice_Success(t *testing.T) {
	mockStorage := NewMockOfficeStorage()
	usecase := commands.NewRegisterOfficeUsecase(mockStorage)

	params := commands.RegisterOfficeParams{
		Number:      101,
		LeasePeriod: 12,
		Status:      "Active",
	}

	err := usecase.Handle(params)

	assert.NoError(t, err)
}

func TestRegisterOffice_Duplicate(t *testing.T) {
	mockStorage := NewMockOfficeStorage()
	usecase := commands.NewRegisterOfficeUsecase(mockStorage)

	existingOffice, _ := office.NewOffice(101, 12, "Active")
	_ = mockStorage.Save(existingOffice)

	params := commands.RegisterOfficeParams{
		Number:      101,
		LeasePeriod: 6,
		Status:      "Inactive",
	}

	err := usecase.Handle(params)

	assert.Error(t, err)
	assert.Equal(t, office.ErrOfficeAlreadyExists, err)
}

func TestRegisterOffice_InvalidNumber(t *testing.T) {
	mockStorage := NewMockOfficeStorage()
	usecase := commands.NewRegisterOfficeUsecase(mockStorage)

	params := commands.RegisterOfficeParams{
		Number:      -5, // Número inválido
		LeasePeriod: 12,
		Status:      "Active",
	}

	err := usecase.Handle(params)

	assert.Error(t, err)
}

func TestRegisterOffice_InvalidLeasePeriod(t *testing.T) {
	mockStorage := NewMockOfficeStorage()
	usecase := commands.NewRegisterOfficeUsecase(mockStorage)

	params := commands.RegisterOfficeParams{
		Number:      101,
		LeasePeriod: -3,
		Status:      "Active",
	}

	err := usecase.Handle(params)

	assert.Error(t, err)
}

func TestRegisterOffice_InvalidStatus(t *testing.T) {
	mockStorage := NewMockOfficeStorage()
	usecase := commands.NewRegisterOfficeUsecase(mockStorage)

	params := commands.RegisterOfficeParams{
		Number:      101,
		LeasePeriod: 12,
		Status:      "Unknown",
	}

	err := usecase.Handle(params)

	assert.Error(t, err)
}
