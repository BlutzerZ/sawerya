package dto

type CreateTransactionRequest struct {
	TransactionID   string `json:"transactionID" binding:"required"`
	Amount          int    `json:"amount" binding:"required"`
	Description     string `json:"description"`
	InvoiceDuration int    `json:"invoiceDuration" binding:"required"`
	Sender          string `json:"sender" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Type            uint8  `json:"type" binding:"required"`
	PaymentMethod string 
}
