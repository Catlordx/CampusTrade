package user

import (
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/permission"
	"github.com/Catlordx/CampusTrade/internal/service/operation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InquireAnyoneInfo
//
//	@Description: 查询其他用户信息，返回信息包括用户手机号和用户角色
//	@param	c	gin.Context
func InquireAnyoneInfo(c *gin.Context) {
	username := c.PostForm("username")
	inquiredUsername := c.PostForm("inquired_username")

	user := operation.User(db, username)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未找到用户信息",
		})
		return
	}

	if operation.HasPermission(db, username, permission.InquireAnyoneInfo) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有查询其他用户信息的权限",
		})
		return
	}

	inquiredUser := operation.User(db, inquiredUsername)
	if inquiredUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未找到被修改用户的信息",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage":      "查询用户信息成功",
		"phone_number": inquiredUser.PhoneNumber,
		"role":         inquiredUser.Role,
	})
}

// ModifyAnyoneInfo
//
//	@Description: 修改其他用户信息，
//	@param	c	c.ginContext
func ModifyAnyoneInfo(c *gin.Context) {
	username := c.PostForm("username")
	modifiedUser := c.PostForm("modified_user")
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

	if operation.HasPermission(db, username, permission.ModifyAnyoneInfo) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改用户信息的权限",
		})
		return
	}

	if operation.ModifyUser(db, modifiedUser, newUsername, newRealName, newPassword, newPhoneNumber) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未找到被修改用户的信息",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "修改用户信息成功",
	})
}
