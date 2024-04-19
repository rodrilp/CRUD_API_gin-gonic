package main

import (
	"crud_api_gin-gonic/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	userGroup := r.Group("/user")
	userGroup.POST("", controllers.CreateUser)
	userGroup.GET("/:id", controllers.GetUserById)
	userGroup.GET("", controllers.GetUsers)
	userGroup.DELETE("/:id", controllers.DeleteUser)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	return r
}
func main() {
	r := setupRouter()

	r.Run()
}
