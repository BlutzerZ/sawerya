package routes

import (
	"blutzerz/sawerya/api/controllers"
	"blutzerz/sawerya/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(group *gin.RouterGroup) {
	ac := controllers.NewAuthController()

	group.POST("/login", ac.LoginUser)
	group.POST("/refresh-token", middleware.JWTAuth(), ac.RefreshToken)
}
