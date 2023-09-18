package routes

import (
	"blutzerz/sawerya/api/controllers"
	"blutzerz/sawerya/middleware"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(group *gin.RouterGroup) {
	tc := controllers.NewTransactionController()

	group.POST("/payment", middleware.JWTAuth(), tc.CreatePayment)
	group.POST("/payment/callback", tc.PaymentCallback)
}
