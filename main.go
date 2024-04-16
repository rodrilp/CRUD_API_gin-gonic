package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	return r
}
func main() {
	r := setupRouter()

	r.Run()
}
