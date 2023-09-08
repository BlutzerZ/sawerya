package service

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/api/repository"
	"fmt"
	"net/http"
)

type TransactionService struct {
	repository *repository.TransactionRepository
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		repository: repository.NewRepositoryService(),
	}
}

func (s *TransactionService) CreateInvoice(TID int, req *dto.CreateTransactionRequest) {
	reqUrl := fmt.Sprintf("https://api.xendit.co/v2/invoices")

	jsonBody := []byte(`{
		"external_id": TID,
		"amount": req.Amount,
		"amount": req.Amount,
	}`)
	req, err := http.NewRequest(http.MethodPost, reqUrl, body)
}

func (s *TransactionService) CreateTransaction {

}
