package service

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/configs"
	"fmt"
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

func TestAlertService(t *testing.T) {
	configs.InitDB()
	s := NewAlertService()

	t.Run("GetAlertByUserID", func(t *testing.T) {
		result, err := s.GetAlertByUserID(1)
		fmt.Println(result)
		assert.Equal(t, nil, err, "must be nil")
	})

	t.Run("EditAlert", func(t *testing.T) {
		req := new(dto.UpdateAlertRequest)
		req.EnableGif = 0
		req.Sound = "xtingtongsss"

		err := s.UpdateAlert(8, req)
		assert.Equal(t, nil, err, "must be nil")
	})
	t.Run("EditAlertDesign", func(t *testing.T) {
		req := new(dto.UpdateAlertDesignRequest)
		req.Duration = 10
		req.TextColor = "#21d292"

		err := s.UpdateAlertDesign(1, req)
		assert.Equal(t, nil, err, "must be nil")
	})
}

func TestTransactionService(t *testing.T) {
	configs.InitDB()
	s := NewTransactionService()

	req := new(dto.CreateTransactionRequest)
	req.Amount = 5000
	req.Email = "pler@gmail.com"
	req.InvoiceDuration = 10
	req.Type = 1
	req.Sender = "jamal"

	transaction := models.Transaction{
		Amount:        req.Amount,
		Description:   req.Description,
		Sender:        req.Sender,
		Email:         req.Email,
		PaymentMethod: "",
		TypeID:        req.Type,
	}

	err := s.CreateTransaction(transaction)
	assert.Equal(t, nil, err, "must be nil")

}
