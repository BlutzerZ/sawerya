package routes

import (
	"blutzerz/sawerya/api/controllers"
	"blutzerz/sawerya/middleware"

	"github.com/gin-gonic/gin"
)

func OverlayRoutes(group *gin.RouterGroup) {
	ac := controllers.NewAlertController()

	group.GET("/overlay/alert/:id", middleware.JWTAuth(), ac.GetAlertByUserID)
	group.PUT("/overlay/alert", middleware.JWTAuth(), ac.UpdateAlert)
	group.PUT("/overlay/alert/design", middleware.JWTAuth(), ac.UpdateAlertDesign)

	// r.GET("/overlay/mediashare", controllers.GetOverlayAlertUrl)
	// r.PUT("/overlay/mediashare", controllers.EditOverlayAlert)

	// r.GET("/overlay/subathon", controllers.GetOverlayAlertUrl)
	// r.PUT("/overlay/subathon", controllers.EditOverlayAlert)

	// r.GET("/overlay/qrcode", controllers.GetOverlayAlertUrl)
	// r.PUT("/overlay/qrcode", controllers.EditOverlayAlert)

	// r.GET("/overlay/qrcode", controllers.GetOverlayAlertUrl)
	// r.PUT("/overlay/qrcode", controllers.EditOverlayAlert)
}
