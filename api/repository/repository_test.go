package repository

import (
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/configs"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
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
		_, result := ur.FindByID(1)
		assert.Equal(t, nil, result, "result must be nil")
	})
	t.Run("updateUsername", func(t *testing.T) {
		result := ur.Update(1, "username", "testnew")
		assert.Equal(t, nil, result, "result must be nil")
	})
	t.Run("updatePassword", func(t *testing.T) {
		result := ur.Update(1, "password", "passwordnew")
		assert.Equal(t, nil, result, "result must be nil")
	})
	t.Run("delete", func(t *testing.T) {
		result := ur.Delete(1)
		assert.Equal(t, nil, result, "result must be nil")

	})
}

func TestAlertRepository(t *testing.T) {
	configs.InitDB()

	ar := NewAlertRepository()

	alert := models.Alert{
		EnableGif:       0,
		MinAmountNotify: 9000,
		MinAmountGIF:    5000,
		Sound:           "default",
	}
	alertDesign := models.AlertDesign{
		BackgroundColor: "#199999",
		HighlightColor:  "#000000",
		TextColor:       "#000000",
		TextTemplate:    "baru saja memberikan",
		Border:          0,
		TextTickness:    100,
		Duration:        5,
		Font:            "arial",
	}

	t.Run("updateALert", func(t *testing.T) {
		result := ar.UpdateAlert(8, &alert)
		assert.Equal(t, nil, result, "result must be nil")
	})

	var alertResult models.Alert
	t.Run("getAlertByID", func(t *testing.T) {
		var result error
		alertResult, result = ar.FindAlertByUserID(1)
		assert.Equal(t, nil, result, "result mus be nil")
	})
	alertResult.AlertDesign = &alertDesign

	fmt.Println(alertResult.ID)

	t.Run("updateAlertDesign", func(t *testing.T) {
		result := ar.UpdateAlertDesign(&alertResult)
		assert.Equal(t, nil, result, "result must be nil")
	})
}
