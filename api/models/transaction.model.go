package models

import (
	"fmt"
	"math/rand"

	"gorm.io/gorm"
)

type Transaction struct {
	ID              uint `gorm:"type:integer; primaryKey"`
	Amount          uint
	Description     string
	Sender          string
	Email           string
	PaymentMethod   string
	TypeID          uint
	TransactionType TransactionType `gorm:"foreignKey:TypeID"`
}

type TransactionType struct {
	ID   uint `gorm:"primaryKey"`
	Type string
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) {
	randTID := fmt.Sprintf("sawerya-%s", randStringRune(8))
	tx.Statement.SetColumn("ID", randTID)
}

func randStringRune(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
