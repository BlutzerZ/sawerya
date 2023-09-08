package models

import (
	"fmt"
	"math/rand"

	"gorm.io/gorm"
)

type Transaction struct {
	ID              string
	Amount          int
	Description     string
	InvoiceDuration int
	Sender          string
	Email           string
	PaymentMethod   string
	TypeID          uint8
	TransactionType TransactionType `gorm:"foreignKey:TypeID"`
}

type TransactionType struct {
	ID   uint
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
