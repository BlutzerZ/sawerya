package routes

import (
	"blutzerz/sawerya/api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(group *gin.RouterGroup) {
	ac := controllers.NewAuthController()

	group.POST("/login", ac.LoginUser)
	// group.PUT("/logout", ac.Logout)
}
