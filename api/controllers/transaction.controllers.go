package controllers

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/api/service"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TransactionController struct {
	service     *service.TransactionService
	serviceUser *service.UserService
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		service:     service.NewTransactionService(),
		serviceUser: service.NewUserService(),
	}
}

func (tc *TransactionController) CreatePayment(c *gin.Context) {
	// REQ BODY
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

	// Service
	user, err := tc.serviceUser.GetUserByUsername(req.Receiver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	transaction, err := tc.service.CreateTransaction(user.ID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	fmt.Println(transaction)
	resp, err := tc.service.CreateInvoice(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})

}

func (tc *TransactionController) PaymentCallback(c *gin.Context) {
	// Validate Header
	callbackToken := c.GetHeader("x-callback-token")
	shouldBeCallbackToken := os.Getenv("XENDIT_CALLBACK_TOKEN")
	if callbackToken != shouldBeCallbackToken {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "invalid callback token",
		})
		return
	}

	req := new(dto.InvoiceCallbackRequest)
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

	err = tc.service.UpdateTransaction(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "callback received",
	})
}
