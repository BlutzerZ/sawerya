package dto

type CreateTransactionRequest struct {
	Amount      uint   `json:"amount" binding:"required"`
	Description string `json:"description"`
	Sender      string `json:"sender" binding:"required"`
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

type InvoiceCallbackResponse struct {
	ExternalID    string `json:"external_id"`
	PaymentMethod string `json:"payment_method"`
	PaidAmount    string `json:"paid_amount"`
	PaidAt        string `json:"paid_at"`
	Created       string `json:"created"`
	Updated       string `json:"updated"`
}
