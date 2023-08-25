package repository

import (
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/configs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	configs.InitDB()

	ur := NewUserRepository()

	user := models.User{
		Username: "test",
		Email:    "test@email.com",
		Password: "test",
	}

	t.Run("create", func(t *testing.T) {
		result := ur.Create(user)
		assert.Equal(t, nil, result, "result must be nil")
	})
	t.Run("find", func(t *testing.T) {
		result := ur.FindByID(1)
		assert.Equal(t, nil, result, "result must be nil")
	})
	t.Run("update", func(t *testing.T) {
		result := ur.Update(1, "username", "testnew")
		assert.Equal(t, nil, result, "result must be nil")
	})
	t.Run("delete", func(t *testing.T) {
		result := ur.Delete(1)
		assert.Equal(t, nil, result, "result must be nil")

	})
}
