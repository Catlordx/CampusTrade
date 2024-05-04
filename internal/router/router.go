package router

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	setProductRouter(r)
	setUserRouter(r)
	setOrderRouter(r)
	r.GET("/hello", helloHandler)
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}
