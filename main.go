package main

import (
	controller "beautyproject/services/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.RegisterUserController(r)
	r.Run()
}
