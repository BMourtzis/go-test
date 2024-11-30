package controller

import (
	"beautyproject/services/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserController(router *gin.Engine) {
	v1 := router.Group("/v1/user")
	{
		v1.GET("/", get)
		v1.GET("/:id", getById)
		v1.POST("/create", create)
	}
}

func get(c *gin.Context) {
	c.JSON(http.StatusOK, models.NewUser())
}

func getById(c *gin.Context) {
	id := c.Param("id")

	user, err := models.NewUserById(id)

	if err != nil {
		c.String(http.StatusBadRequest, "")
	}

	c.JSON(http.StatusOK, user)
}

func create(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	//TODO: add call to service

	c.JSON(http.StatusOK, user)
}
