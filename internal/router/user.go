package router

import (
	"github.com/Catlordx/CampusTrade/internal/service/user"
	"github.com/gin-gonic/gin"
)

func setUserRouter(r *gin.Engine) {
	userGroup := r.Group("user")
	{
		// TODO 用户注册
		userGroup.POST("/register", user.Register)
		// TODO 用户登录
		userGroup.POST("/login", user.Login)
		// TODO 获得用户信息
		userGroup.GET("/profile/info", user.InquireInfo)
		userGroup.GET("/profile/info/anyone_info", user.InquireAnyoneInfo)
		// TODO 修改用户信息
		userGroup.PUT("/profile/modify_username", user.ModifyUsername)
		userGroup.PUT("/profile/modify_real_name", user.ModifyRealName)
		userGroup.PUT("/profile/modify_password", user.ModifyPassword)
		userGroup.PUT("/profile/modify_phone_number", user.ModifyPhoneNumber)
		userGroup.PUT("/profile/modify_role", user.ModifyRole)
		userGroup.PUT("/profile/modify_anyone_username", user.ModifyAnyoneUsername)
		userGroup.PUT("/profile/modify_anyone_real_name", user.ModifyAnyoneRealName)
		userGroup.PUT("/profile/modify_anyone_password", user.ModifyAnyonePassword)
		userGroup.PUT("/profile/modify_anyone_phone_number", user.ModifyAnyonePhoneNumber)
		userGroup.PUT("/profile/modify_anyone_role", user.ModifyAnyoneRole)
	}
}
