package routes

import (
	"blutzerz/sawerya/api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(group *gin.RouterGroup) {
	uc := controllers.NewUserController()

	// group.GET("/user", controllers.GetAllUser)
	// r.GET("/user/{id}", controllers.GetUserByID)
	group.POST("/user", uc.CreateNewUser)
	group.PUT("/user/username", uc.UpdateUsername)
	group.PUT("/user/password", uc.UpdatePassword)
}
