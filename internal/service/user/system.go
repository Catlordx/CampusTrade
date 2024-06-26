package user

import (
	"github.com/Catlordx/CampusTrade/internal/core/config"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/role"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/status"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/user"
	"github.com/Catlordx/CampusTrade/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// Register
//
//	@Description: 注册用户
//	@param	c	gin.Context
func Register(c *gin.Context) {
	var _userInfo userInfo
	err := c.ShouldBind(&_userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	appContext := c.MustGet("appContext").(*config.AppContext)
	//注册默认为user角色
	err = user.AddUser(appContext.DB, _userInfo.Username, "", _userInfo.Password, "", role.User)
	if err != nil {
		switch err.Error() {
		case "user already exists":
			c.JSON(http.StatusBadRequest, gin.H{"message": "用户已经存在"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
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
	var _userInfo userInfo
	err := c.ShouldBind(&_userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	appContext := c.MustGet("appContext").(*config.AppContext)

	_user := user.GetUserByUsername(appContext.DB, _userInfo.Username)
	if _user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不存在",
		})
		return
	}
	if user.CheckPassword(_user.Password, _userInfo.Password) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码错误",
		})
		return
	}
	if _user.Status == status.Online {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户已经登录",
		})
		return
	}

	_ = user.ModifyStatus(appContext.DB, _user.ID, status.Online)
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
