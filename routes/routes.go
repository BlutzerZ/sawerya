package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")

	UserRoutes(apiGroup)
	AuthRoutes(apiGroup)

	r.Run()
}
