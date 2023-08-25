package main

import (
	"blutzerz/sawerya/configs"
	"blutzerz/sawerya/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	configs.InitDB()

	r := gin.Default()

	routes.RegisterRoutes(r)
}
