package router

import (
	"github.com/Catlordx/CampusTrade/internal/middleware"
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

		userGroup.Use(middleware.Authenticate())
		{
			// TODO 获得用户信息
			userGroup.GET("/profile/info", user.InquireInfo)
			userGroup.GET("/profile/anyone_info", user.InquireAnyoneInfo)

			// TODO 修改用户信息
			userGroup.PUT("/profile/modify_info", user.ModifyInfo)
			userGroup.PUT("/profile/modify_anyone_role", user.ModifyAnyoneInfo)

			// TODO 商品收藏管理
			userGroup.POST("/favorite/add", user.AddFavorite)
			userGroup.DELETE("/favorite/remove", user.RemoveFavorite)
			userGroup.GET("/favorite/list/:sort/:reverse/:page/:count", user.FavoriteList)
		}
	}
}
