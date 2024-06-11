package user

import (
	"github.com/Catlordx/CampusTrade/internal/core/config"
	_ "github.com/Catlordx/CampusTrade/internal/db/mysql"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/permission"
	_user "github.com/Catlordx/CampusTrade/internal/db/mysql/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InquireAnyoneInfo
//
//	@Description: 查询其他用户信息，返回信息包括用户手机号和用户角色
//	@param	c	*gin.Context
func InquireAnyoneInfo(c *gin.Context) {
	user := GetUserFromClaims(c)
	if user == nil {
		return
	}

	if CheckUserPermission(c, user, permission.InquireAnyoneInfo) == false {
		return
	}

	inquiredUsername := c.Query("inquired_username")

	appContext := c.MustGet("appContext").(*config.AppContext)
	inquiredUser := _user.GetUserByUsername(appContext.DB, inquiredUsername)
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
//	@param	c	*c.ginContext
func ModifyAnyoneInfo(c *gin.Context) {
	user := GetUserFromClaims(c)
	if user == nil {
		return
	}

	if CheckUserPermission(c, user, permission.ModifyAnyoneInfo) == false {
		return
	}

	modifiedUserUsername := c.PostForm("modified_user")
	newUsername := c.PostForm("new_username")
	newRealName := c.PostForm("new_real_name")
	newPassword := c.PostForm("new_password")
	newPhoneNumber := c.PostForm("new_phone_number")

	appContext := c.MustGet("appContext").(*config.AppContext)
	modifiedUser := _user.GetUserByUsername(appContext.DB, modifiedUserUsername)
	if modifiedUser == nil {
		c.JSON(http.StatusOK, gin.H{"message": "未找到被修改用户的信息"})
		return
	}

	CheckModifyUser(c, modifiedUser.ID, newUsername, newRealName, newPassword, newPhoneNumber)
}
