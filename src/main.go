package main

import (
	controller "beautyproject/services/src/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	controller.RegisterUserController(r)
	r.Run()
}
