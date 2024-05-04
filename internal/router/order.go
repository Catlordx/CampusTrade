package router

import "github.com/gin-gonic/gin"

func setOrderRouter(r *gin.Engine) {
	orderGroup := r.Group("order")
	{
		orderGroup.POST("/")
		orderGroup.GET("/:orderId")
		orderGroup.GET("/user/:userId")
		orderGroup.PUT("/:orderId/complete")
		orderGroup.PUT("/:orderId/cancel")
	}
}
