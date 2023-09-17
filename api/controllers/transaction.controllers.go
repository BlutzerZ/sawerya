package controllers

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/api/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TransactionController struct {
	service *service.TransactionService
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		service: service.NewTransactionService(),
	}
}

func (tc *TransactionController) CreatePayment(c *gin.Context) {
	req := new(dto.CreateTransactionRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, "invalid request")
			return
		}
		var errorMessage string
		for _, e := range validationErrs {
			errorMessage = fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			break
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	transaction, err := tc.service.CreateTransaction(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	fmt.Println(transaction)
	resp, err := tc.service.CreateInvoice(&transaction)
	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})

	return
}
