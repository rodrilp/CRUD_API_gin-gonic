package main

import (
	"crud_api_gin-gonic/config"
	"crud_api_gin-gonic/models"
)

func main()  {
	config.DatabaseConnection().AutoMigrate(&models.User{})
}