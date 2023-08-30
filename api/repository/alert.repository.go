package repository

import (
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/configs"

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

func (ar *AlertRepository) FindAlertByUserID(userID int) (models.Alert, error) {
	var alert models.Alert

	err := ar.DB.Where("user_id = ?", userID).Preload("AlertDesign").First(&alert).Error
	return alert, err
}

func (ar *AlertRepository) UpdateAlert(ID uint, alert *models.Alert) error {
	tx := ar.DB.Begin()

	err := tx.Model(&models.Alert{}).Where("id = ?", ID).Updates(&alert).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (ar *AlertRepository) UpdateAlertDesign(alert *models.Alert) error {
	tx := ar.DB.Begin()

	err := tx.Where("alert_id = ?", alert.ID).Updates(&alert.AlertDesign).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
