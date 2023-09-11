package dto

type CreateTransactionRequest struct {
	TransactionID   string `json:"transactionID" binding:"required"`
	Amount          uint   `json:"amount" binding:"required"`
	Description     string `json:"description"`
	InvoiceDuration uint   `json:"invoiceDuration" binding:"required"`
	Sender          string `json:"sender" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Type            uint   `json:"type" binding:"required"`
	PaymentMethod   string
}
