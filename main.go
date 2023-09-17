package main

import (
	"blutzerz/sawerya/configs"
	"blutzerz/sawerya/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	configs.InitDB()
	r := gin.Default()
	routes.RegisterRoutes(r)
}
