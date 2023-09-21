package models

type TransactionType struct {
	ID   uint `gorm:"primaryKey"`
	Type string
}
