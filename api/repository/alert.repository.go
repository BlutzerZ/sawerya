package repository

import (
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/configs"
	"fmt"

	"gorm.io/gorm"
)

type AlertRepository struct {
	DB *gorm.DB
}

func NewAlertRepository() *AlertRepository {
	DB := configs.GetDB()
	return &AlertRepository{
		DB: DB,
	}
}

func (ar *AlertRepository) FindAlertByUserID(userID uint) (models.Alert, error) {
	var alert models.Alert

	err := ar.DB.Where("user_id = ?", userID).Preload("AlertDesign").First(&alert).Error
	return alert, err
}

func (ar *AlertRepository) UpdateAlert(ID uint, alert *models.Alert) error {
	tx := ar.DB.Begin()

	fmt.Println(alert)

	err := tx.Model(&models.Alert{}).Where("user_id = ?", ID).Updates(map[string]interface{}{
		"enable_gif":        alert.EnableGif,
		"min_amount_notify": alert.MinAmountNotify,
		"min_amount_gif":    alert.MinAmountGIF,
		"sound":             alert.Sound,
		"word_filter":       alert.WordFilter,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (ar *AlertRepository) UpdateAlertDesign(alert *models.Alert) error {
	tx := ar.DB.Begin()

	err := tx.Model(&models.AlertDesign{}).Where("alert_id = ?", alert.ID).Updates(map[string]interface{}{
		"background_color": alert.AlertDesign.BackgroundColor,
		"highlight_color":  alert.AlertDesign.HighlightColor,
		"text_color":       alert.AlertDesign.TextColor,
		"text_template":    alert.AlertDesign.TextTemplate,
		"border":           alert.AlertDesign.Border,
		"text_tickness":    alert.AlertDesign.TextTickness,
		"duration":         alert.AlertDesign.Duration,
		"font":             alert.AlertDesign.Font,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
