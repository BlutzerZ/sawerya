package dto

import "time"

type CreateTransactionRequest struct {
	Amount      uint   `json:"amount" binding:"required"`
	Description string `json:"description"`
	Sender      string `json:"sender" binding:"required"`
	Receiver    string `json:"receiver" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Type        uint   `json:"type" binding:"required"`
}

type CreateTransactionResponse struct {
	ExternalID string `json:"external_id"`
	Amount     uint   `json:"amount"`
	Currency   string `json:"currency"`
	InvoiceUrl string `json:"invoice_url"`
	ExpiredAt  string `json:"expiry_date"`
	CreatedAt  string `json:"created"`
	UpdatedAt  string `json:"updated"`
}

type InvoiceCallbackRequest struct {
	ExternalID    string    `json:"external_id"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	PaidAmount    uint      `json:"paid_amount"`
	PaidAt        time.Time `json:"paid_at"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
}
