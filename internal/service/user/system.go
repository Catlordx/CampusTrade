package user

import (
	"github.com/Catlordx/CampusTrade/internal/core/config"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/user"
	"github.com/Catlordx/CampusTrade/internal/utils"
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

	appContext := c.MustGet("appContext").(*config.AppContext)

	if user.AddUser(appContext.DB, username, realName, password, phoneNumber, role) == false {
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

	appContext := c.MustGet("appContext").(*config.AppContext)

	_user := user.GetUserByUsername(appContext.DB, username)
	if _user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不存在",
		})
		return
	}
	if password != _user.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码错误",
		})
		return
	}

	token, err := utils.GenerateToken(_user.ID, _user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
	})
}
