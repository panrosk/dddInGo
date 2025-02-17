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

type MockOfficeStorage struct {
	offices []*entities.Office
}

func NewMockOfficeStorage() *MockOfficeStorage {
	return &MockOfficeStorage{offices: make([]*entities.Office, 0)}
}

func (m *MockOfficeStorage) Save(o *entities.Office) error {
	m.offices = append(m.offices, o)
	return nil
}

func (m *MockOfficeStorage) FindById(id any) (*entities.Office, error) {
	for _, o := range m.offices {
		if o.GetOffice()["id"] == id {
			return o, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *MockOfficeStorage) FindAll() ([]*entities.Office, error) {
	return m.offices, nil
}

func (m *MockOfficeStorage) FindByFilter(filterFunc func(*entities.Office) bool) ([]*entities.Office, error) {
	var result []*entities.Office
	for _, o := range m.offices {
		if filterFunc(o) {
			result = append(result, o)
		}
	}
	return result, nil
}

// Verificación explícita de implementación de la interfaz
var _ ports.RepositoryPort[*entities.Office] = (*MockOfficeStorage)(nil)

func TestRegisterOffice_Success(t *testing.T) {
	mockStorage := &MockOfficeStorage{}
	usecase := commands.NewRegisterOfficeUsecase(mockStorage)

	params := commands.RegisterOfficeParams{
		Number:      101,
		LeasePeriod: 12,
		Status:      "Active",
	}

	office, err := usecase.Execute(params)

	assert.NoError(t, err)
	assert.NotNil(t, office)
	assert.Equal(t, params.Number, office.GetOffice()["number"])
	assert.Equal(t, params.LeasePeriod, office.GetOffice()["lease_period"])
	assert.Equal(t, params.Status, office.GetOffice()["status"])
}

func TestRegisterOffice_Duplicate(t *testing.T) {
	mockStorage := &MockOfficeStorage{}
	usecase := commands.NewRegisterOfficeUsecase(mockStorage)

	existingOffice, _ := entities.NewOffice(101, 12, "Active")
	mockStorage.Save(existingOffice)

	params := commands.RegisterOfficeParams{
		Number:      101,
		LeasePeriod: 6,
		Status:      "Inactive",
	}

	office, err := usecase.Execute(params)

	assert.Error(t, err)
	assert.Nil(t, office)
	assert.Equal(t, domain_errors.ErrOfficeAlreadyExists, err)
}

func TestRegisterOffice_InvalidNumber(t *testing.T) {
	mockStorage := &MockOfficeStorage{}
	usecase := commands.NewRegisterOfficeUsecase(mockStorage)

	params := commands.RegisterOfficeParams{
		Number:      -5, // Número inválido
		LeasePeriod: 12,
		Status:      "Active",
	}

	office, err := usecase.Execute(params)

	assert.Error(t, err)
	assert.Nil(t, office)
}

func TestRegisterOffice_InvalidLeasePeriod(t *testing.T) {
	mockStorage := &MockOfficeStorage{}
	usecase := commands.NewRegisterOfficeUsecase(mockStorage)

	params := commands.RegisterOfficeParams{
		Number:      101,
		LeasePeriod: -3, // LeasePeriod inválido
		Status:      "Active",
	}

	office, err := usecase.Execute(params)

	assert.Error(t, err)
	assert.Nil(t, office)
}

func TestRegisterOffice_InvalidStatus(t *testing.T) {
	mockStorage := &MockOfficeStorage{}
	usecase := commands.NewRegisterOfficeUsecase(mockStorage)

	params := commands.RegisterOfficeParams{
		Number:      101,
		LeasePeriod: 12,
		Status:      "Unknown", // Estado inválido
	}

	office, err := usecase.Execute(params)

	assert.Error(t, err)
	assert.Nil(t, office)
}
