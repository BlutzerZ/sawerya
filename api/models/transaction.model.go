package models

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID              string `gorm:"type:string; primaryKey"`
	Amount          uint
	Description     string
	Sender          string
	Email           string
	PaymentMethod   string
	Status          string
	PaidAt          time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	TypeID          uint
	TransactionType TransactionType `gorm:"foreignKey:TypeID"`
}

type TransactionType struct {
	ID   uint `gorm:"primaryKey"`
	Type string
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	randTID := fmt.Sprintf("sawerya-%s", randStringRune(16))

	tx.Statement.SetColumn("ID", randTID)
	tx.Statement.SetColumn("Status", "UNPAID")

	return nil

}

func randStringRune(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
