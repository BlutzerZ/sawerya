package service

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/api/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type TransactionService struct {
	repository *repository.TransactionRepository
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		repository: repository.NewTransactionRepository(),
	}
}

func (s *TransactionService) CreateInvoice(transaction *models.Transaction) (*dto.CreateTransactionResponse, error) {

	url := "https://api.xendit.co/v2/invoices"

	// BODY
	data := map[string]interface{}{
		"external_id": transaction.ID,
		"amount":      transaction.Amount,
		"currency":    "IDR",
		"customer": map[string]interface{}{
			"given_name": transaction.Sender,
			"email":      transaction.Email,
		},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// HEADER
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Basic "+os.Getenv("XENDIT_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err

	}
	defer resp.Body.Close()

	// RESPONSE BODY
	var createTransactionResponse *dto.CreateTransactionResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&createTransactionResponse); err != nil {
		fmt.Println("Error:", err)
		return nil, err

	}

	return createTransactionResponse, nil
}

func (s *TransactionService) CreateTransaction(receiverID uint, req *dto.CreateTransactionRequest) (models.Transaction, error) {

	transaction := models.Transaction{
		Amount:      req.Amount,
		Description: req.Description,
		Sender:      req.Sender,
		ReceiverID:  receiverID,
		Email:       req.Email,
		TypeID:      req.Type,
	}

	err := s.repository.CreateNewTransaction(&transaction)
	return transaction, err

}

func (s *TransactionService) UpdateTransaction(req *dto.InvoiceCallbackRequest) error {
	transaction := models.Transaction{
		ID:            req.ExternalID,
		Status:        req.Status,
		PaymentMethod: req.PaymentMethod,
		PaidAt:        req.PaidAt,
		UpdatedAt:     req.Updated,
	}

	err := s.repository.UpdateTransaction(&transaction)

	return err
}
