package controller

import (
	"beautyproject/services/src/db"
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
	users, err := db.ReadUsers()

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, users)
}

func getById(c *gin.Context) {
	id := c.Param("id")

	user, err := models.NewUserById(id)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, user)
}

func create(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	id, err := db.CreateUser(user)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, id)
}
