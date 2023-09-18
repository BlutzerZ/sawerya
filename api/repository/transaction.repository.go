package repository

import (
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/configs"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		db: configs.GetDB(),
	}
}

func (r *TransactionRepository) CreateNewTransaction(transaction *models.Transaction) error {
	err := r.db.Create(&transaction).Error

	return err
}

func (r *TransactionRepository) UpdateTransaction(transaction *models.Transaction) error {
	tx := r.db.Begin()

	err := tx.Model(&models.Transaction{}).Where("ID = ?", transaction.ID).Updates(&transaction).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
