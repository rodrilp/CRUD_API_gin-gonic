package controllers

import (
	"crud_api_gin-gonic/config"
	"crud_api_gin-gonic/models"

	"github.com/gin-gonic/gin"
)
var user models.User
var users []models.User

func CreateUser(c *gin.Context) {
	// Obtain data from the body
	c.Bind(&user)
	
	// Create the user
	result := config.DatabaseConnection().Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func GetUsers(c *gin.Context) {
	// Obtain all the users
	result := config.DatabaseConnection().Find(&users)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"user_list": users})
}

func GetUserById(c *gin.Context) {
	// Obtain data from the uri
	id := c.Param("id")

	// Obtain the user with the id passed as a param
	result := config.DatabaseConnection().First(&user, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"user": user})
}