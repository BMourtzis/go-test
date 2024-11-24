package main

import (
	"beautyproject/services/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", root)
	r.GET("/test", testWithParametets)
	r.Run()
}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, models.NewUser())
}

func testWithParametets(c *gin.Context) {
	age, err := strconv.Atoi(c.DefaultQuery("age", "0"))
	if err != nil {
		age = 0
	}
	firstName := c.DefaultQuery("firstname", "Bill")
	lastName := c.DefaultQuery("lastname", "Mourtzis")

	c.JSON(http.StatusOK, models.NewUserWithParameters(firstName, lastName, age))
}
