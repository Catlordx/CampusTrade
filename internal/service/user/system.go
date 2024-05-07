package user

import (
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
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

	db, dbErr := mysql.DB()
	if dbErr != nil {
		switch dbErr.Error() {
		case "failed to load config":
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "加载配置失败",
			})
		case "failed to connect database":
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "数据库连接失败",
			})
		}
		return
	}

	_, err := operation.AddUser(db, username, realName, password, phoneNumber, role)
	if err != nil {
		switch err.Error() {
		case "username is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入用户名",
			})
		case "user already exists":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "用户已经存在",
			})
		case "username is too short":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "用户名过短",
			})
		case "username is invalid":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "用户名包含特殊字符",
			})
		case "real name is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入真实名字",
			})
		case "real name is invalid":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "真实名字包含特殊字符",
			})
		case "password is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入密码",
			})
		case "password is too short":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "密码过短",
			})
		case "password is invalid":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "密码包含特殊字符",
			})
		case "phone number is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入手机号",
			})
		case "phone number is invalid":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入正确格式的手机号",
			})
		case "role is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入用户角色",
			})
		case "role is not exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "不存在这种用户角色",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未知的错误",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户注册成功",
		})
	}
}

// Login
//
//	@Description: 用户登录
//	@param	c	gin.Context
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	db, dbErr := mysql.DB()
	if dbErr != nil {
		switch dbErr.Error() {
		case "failed to load config":
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "加载配置失败",
			})
		case "failed to connect database":
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "数据库连接失败",
			})
		}
		return
	}

	user := operation.User(db, username)
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
