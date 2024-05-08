package user

import (
	"github.com/Catlordx/CampusTrade/internal/db/mysql/permission"
	"github.com/Catlordx/CampusTrade/internal/service/operation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InquireInfo
//
//	@Description: 查询用户自己的信息，返回信息包括用户的真实名字、手机号和用户角色
//	@param	c	gin.Context
func InquireInfo(c *gin.Context) {
	username := c.PostForm("username")

	user := operation.User(db, username)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未找到用户信息",
		})
		return
	}

	if operation.HasPermission(db, username, permission.InquireInfo) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有查询自己信息的权限",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"massage":      "查询用户信息成功",
		"real_name":    user.RealName,
		"phone_number": user.PhoneNumber,
		"role":         user.Role,
	})
}

func ModifyInfo(c *gin.Context) {
	username := c.PostForm("username")
	newUsername := c.PostForm("new_username")
	newRealName := c.PostForm("new_real_name")
	newPassword := c.PostForm("new_password")
	newPhoneNumber := c.PostForm("new_phone_number")

	user := operation.User(db, username)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未找到用户信息",
		})
		return
	}

	if operation.HasPermission(db, username, permission.ModifyInfo) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改自己信息的权限",
		})
		return
	}

	operation.ModifyUser(db, username, newUsername, newRealName, newPassword, newPhoneNumber)
	c.JSON(http.StatusOK, gin.H{
		"message": "修改用户信息成功",
	})
}
