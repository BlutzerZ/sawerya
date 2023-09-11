package service

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/api/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TransactionService struct {
	repository *repository.TransactionRepository
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		repository: repository.NewTransactionRepository(),
	}
}

func (s *TransactionService) CreateInvoice(TID int, req *dto.CreateTransactionRequest) (string, error) {
	reqUrl := fmt.Sprintf("https://api.xendit.co/v2/invoices")

	requestData := map[string]interface{}{
		"external_id": TID,
		"amount":      req.Amount,
		"customer": map[string]interface{}{
			"email": req.Email,
		},
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	httpReq, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	httpReq.Header.Set("Authorization", "Bearer xnd_public_development_aZho51dSsowWNPOvAs9kgSuyFXAeXHMiWHAVw02v0523s0BHIBPnM0KgGbu1KN")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// READ
	respondBuffer := new(bytes.Buffer)
	respondBuffer.ReadFrom(resp.Body)
	respondString := respondBuffer.String()

	return respondString, nil
}

func (s *TransactionService) CreateTransaction(transaction models.Transaction) error {
	err := s.repository.CreateNewTransaction(transaction)
	return err

}
