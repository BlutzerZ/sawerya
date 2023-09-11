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

func (r *TransactionRepository) CreateNewTransaction(transaction models.Transaction) error {
	err := r.db.Create(transaction).Error

	return err
}
