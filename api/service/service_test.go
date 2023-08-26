package service

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/configs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	configs.InitDB()

	s := NewUserService()

	t.Run("createUser", func(t *testing.T) {
		req := new(dto.RegisterUserRequest)
		req.Email = "test@email.com"
		req.Username = "test"
		req.Password = "testpwd123"
		result := s.CreateUser(req)
		assert.Equal(t, nil, result, "result must be nil")
	})
	t.Run("MatchPwd", func(t *testing.T) {
		pwd := "testpwd123"
		_, result := s.isPasswordMatch(1, pwd)
		assert.Equal(t, nil, result, "result must be nil")

	})

	t.Run("UpdateUsername", func(t *testing.T) {
		req := new(dto.UpdateUsernameRequest)
		req.ID = 1
		req.Username = "test123"
		req.Password = "testpwd123"
		result := s.UpdateUsername(req)
		assert.Equal(t, nil, result, "result must be nil")
	})
	t.Run("UpdatePassword", func(t *testing.T) {
		req := new(dto.UpdatePasswordRequest)
		req.ID = 1
		req.OldPassword = "testpwd123"
		req.Password = "testnewpwd123"
		result := s.UpdatePassword(req)
		assert.Equal(t, nil, result, "result must be nil")
	})
	t.Run("DeleteUser", func(t *testing.T) {
		result := s.DeleteUser(1)
		assert.Equal(t, nil, result, "result must be nil")
	})
}
