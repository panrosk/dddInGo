package entities_test

import (
	"coworking/internal/app/domain/domain_errors"
	"coworking/internal/app/domain/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewHotdesk_Success(t *testing.T) {
	hotdesk, err := entities.NewHotdesk(5)
	assert.NoError(t, err)
	assert.NotNil(t, hotdesk)
	assert.Equal(t, 5, hotdesk.Number.Value())
	assert.Equal(t, "Active", string(hotdesk.GetHotdesk()["status"].(string)))
	assert.NotEmpty(t, hotdesk.GetHotdesk()["id"])
	assert.NotEmpty(t, hotdesk.GetHotdesk()["created_at"])
	assert.NotEmpty(t, hotdesk.GetHotdesk()["updated_at"])
}

func TestNewHotdesk_InvalidNumber(t *testing.T) {
	hotdesk, err := entities.NewHotdesk(-1)

	assert.Error(t, err)
	assert.Nil(t, hotdesk)
	assert.Equal(t, err, domain_errors.ErrInvalidHotDeskNumber)
}

func TestGetHotdesk(t *testing.T) {
	hotdesk, _ := entities.NewHotdesk(10)

	hotdeskData := hotdesk.GetHotdesk()

	assert.Equal(t, hotdesk.Number.Value(), hotdeskData["number"])
	assert.Equal(t, "Active", hotdeskData["status"])
	assert.NotEmpty(t, hotdeskData["id"])

	createdAt, err := time.Parse(time.RFC3339, hotdeskData["created_at"].(string))
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now(), createdAt, time.Second*1)

	updatedAt, err := time.Parse(time.RFC3339, hotdeskData["updated_at"].(string))
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now(), updatedAt, time.Second*1)
}
