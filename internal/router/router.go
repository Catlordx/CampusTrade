package router

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	r.GET("/hello", helloHandler)
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}
