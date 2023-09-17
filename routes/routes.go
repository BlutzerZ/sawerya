package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")

	UserRoutes(apiGroup)
	AuthRoutes(apiGroup)
	OverlayRoutes(apiGroup)
	TransactionRoutes(apiGroup)

	r.Run(":" + os.Getenv("HOST_PORT"))
}
