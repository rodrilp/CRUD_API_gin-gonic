package controllers

import (
	"crud_api_gin-gonic/config"
	"crud_api_gin-gonic/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fails to read the body"})
	}

	user.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fails to read the id"})
	}

	result := config.DatabaseConnection().Model(&models.User{}).Where("id=?", id).Updates(user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	if result.RowsAffected == 0 {
		log.Println("Ninguna fila afectada")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func DeleteUser(c *gin.Context) {
	// Obtain data from the uri
	id := c.Param("id")

	// Obtain the user with the id passed as a param
	result := config.DatabaseConnection().Delete(&user, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	message := fmt.Sprintf("User with ID %s has been delete", id)

	c.JSON(200, gin.H{"message": message})
}
