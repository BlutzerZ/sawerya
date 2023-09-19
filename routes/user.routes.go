package routes

import (
	"blutzerz/sawerya/api/controllers"
	"blutzerz/sawerya/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(group *gin.RouterGroup) {
	uc := controllers.NewUserController()

	group.POST("/user", uc.CreateNewUser)
	group.PUT("/user/username", middleware.JWTAuth(), uc.UpdateUsername)
	group.PUT("/user/password", middleware.JWTAuth(), uc.UpdatePassword)
	group.DELETE("/user", middleware.JWTAuth(), uc.DeleteUser)
}
