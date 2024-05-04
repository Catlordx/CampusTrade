package router

import "github.com/gin-gonic/gin"

func setUserRouter(r *gin.Engine) {
	userGroup := r.Group("user")
	{
		// TODO 用户注册
		userGroup.POST("/register")
		// TODO 用户登录
		userGroup.POST("/login")
		// TODO 获得用户信息
		userGroup.GET("/profile")
		// TODO 修改用户信息
		userGroup.PUT("/profile")
	}
}
