package user

import (
	"github.com/Catlordx/CampusTrade/internal/core"
	"github.com/Catlordx/CampusTrade/internal/service/operation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register
//
//	@Description: 注册用户
//	@param	c	gin.Context
func Register(c *gin.Context) {
	username := c.PostForm("username")
	realName := c.PostForm("real_name")
	password := c.PostForm("password")
	phoneNumber := c.PostForm("phone_number")
	role := c.PostForm("role")

	appContext := c.MustGet("appContext").(*core.AppContext)

	if operation.AddUser(appContext.DB, username, realName, password, phoneNumber, role) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户已经存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "用户注册成功",
	})
}

// Login
//
//	@Description: 用户登录
//	@param	c	gin.Context
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	appContext := c.MustGet("appContext").(*core.AppContext)

	user := operation.User(appContext.DB, username)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不存在",
		})
	} else if password != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码错误",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
		})
	}
}
