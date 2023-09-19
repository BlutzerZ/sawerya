package routes

import (
	"blutzerz/sawerya/api/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(group *gin.RouterGroup) {
	tc := controllers.NewTransactionController()

	group.POST("/payment", tc.CreatePayment)
	group.POST("/payment/callback", tc.PaymentCallback)
}
