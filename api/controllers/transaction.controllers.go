package controllers

import (
	"blutzerz/sawerya/api/dto"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service *service.PaymentService
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		service: service.NewPaymentService,
	}
}

func CreatePayment(c *gin.Context) {
	req := new(dto.CreateInvoiceRequest)

}
