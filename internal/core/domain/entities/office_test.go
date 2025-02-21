package entities_test

import (
	"coworking/internal/app/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOffice_Success(t *testing.T) {
	office, err := entities.NewOffice(101, 12, "Active")

	assert.NoError(t, err)
	assert.NotNil(t, office)
	assert.Equal(t, 101, office.GetOffice()["number"])
	assert.Equal(t, "Active", office.GetOffice()["status"])
	assert.Equal(t, 12, office.GetOffice()["lease_period"])
	assert.NotEmpty(t, office.GetOffice()["id"])
	assert.NotEmpty(t, office.GetOffice()["created_at"])
	assert.NotEmpty(t, office.GetOffice()["updated_at"])
}

func TestNewOffice_DefaultStatus(t *testing.T) {
	office, err := entities.NewOffice(102, 6, "")

	assert.NoError(t, err)
	assert.NotNil(t, office)
	assert.Equal(t, "Active", office.GetOffice()["status"])
}

func TestNewOffice_InvalidNumber(t *testing.T) {
	office, err := entities.NewOffice(-1, 12, "Active")

	assert.Error(t, err)
	assert.Nil(t, office)
}

func TestNewOffice_InvalidLeasePeriod(t *testing.T) {
	office, err := entities.NewOffice(101, -5, "Active")

	assert.Error(t, err)
	assert.Nil(t, office)
}

func TestNewOffice_InvalidStatus(t *testing.T) {
	office, err := entities.NewOffice(101, 12, "Unknown")

	assert.Error(t, err)
	assert.Nil(t, office)
}
