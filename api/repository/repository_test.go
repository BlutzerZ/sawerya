package repository

import (
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/configs"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	err := godotenv.Load("../../.env")
	assert.Equal(t, nil, err, "result must be nil")

	configs.InitDB()

	ur := NewUserRepository()

	user := models.User{
		Username: "test3",
		Email:    "test3@email.com",
		Password: "test3",
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
		user := models.User{
			ID:       1,
			Username: "test3",
		}
		result := ur.Update(&user)
		assert.Equal(t, nil, result, "result must be nil")
	})
	t.Run("updatePassword", func(t *testing.T) {
		user := models.User{
			ID:       1,
			Password: "testpassword3",
		}
		result := ur.Update(&user)
		assert.Equal(t, nil, result, "result must be nil")
	})
	t.Run("delete", func(t *testing.T) {
		result := ur.Delete(35)
		assert.Equal(t, nil, result, "result must be nil")

	})
}

func TestAlertRepository(t *testing.T) {
	err := godotenv.Load("../../.env")
	assert.Equal(t, nil, err, "result must be nil")

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

func TestTransactionRepository(t *testing.T) {
	err := godotenv.Load("../../.env")
	assert.Equal(t, nil, err, "result must be nil")

	configs.InitDB()
	tr := NewTransactionRepository()

	t.Run("createTransaction", func(t *testing.T) {
		transaction := models.Transaction{
			Amount: 5000,
			Sender: "testSender",
			Email:  "test@email.com",
			TypeID: 1,
		}

		err := tr.CreateNewTransaction(&transaction)
		assert.Equal(t, nil, err, "result must be nil")
	})

	t.Run("updateTransaction", func(t *testing.T) {
		tm, err := time.Parse("2006-01-02 15:04:05.000", "2023-09-18 07:43:41.285")
		assert.Equal(t, nil, err, "result must be nil")

		updateTransaction := models.Transaction{
			ID:            "sawerya-Ynmm00fn7ErCq8L5",
			Status:        "PAID",
			PaymentMethod: "BANK_TRANSFER",
			PaidAt:        tm,
		}

		err = tr.UpdateTransaction(&updateTransaction)
		assert.Equal(t, nil, err, "result must be nil")
	})
}
