package user

import (
	"github.com/Catlordx/CampusTrade/internal/core/config"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/permission"
	_user "github.com/Catlordx/CampusTrade/internal/db/mysql/user"
	"github.com/Catlordx/CampusTrade/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InquireInfo
//
//	@Description: 查询用户自己的信息，返回信息包括用户的真实名字、手机号和用户角色
//	@param	c	*gin.Context
func InquireInfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	userClaims, ok := claims.(*utils.CustomClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "claims类型错误",
		})
		return
	}

	appContext := c.MustGet("appContext").(*config.AppContext)
	user := _user.GetUserByID(appContext.DB, userClaims.UserID)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未找到用户信息",
		})
		return
	}

	if _user.HasPermission(appContext.DB, user.Username, permission.InquireInfo) == false {
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

// ModifyInfo
//
//	@Description: 修改自己的用户信息
//	@param	c	*gin.Context
func ModifyInfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	newUsername := c.PostForm("new_username")
	newRealName := c.PostForm("new_real_name")
	newPassword := c.PostForm("new_password")
	newPhoneNumber := c.PostForm("new_phone_number")
	userClaims, ok := claims.(*utils.CustomClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "claims类型错误",
		})
		return
	}

	appContext := c.MustGet("appContext").(*config.AppContext)
	user := _user.GetUserByID(appContext.DB, userClaims.UserID)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未找到用户信息",
		})
		return
	}

	if _user.HasPermission(appContext.DB, user.Username, permission.ModifyInfo) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改自己信息的权限",
		})
		return
	}

	_user.ModifyUser(appContext.DB, user.Username, newUsername, newRealName, newPassword, newPhoneNumber)
	c.JSON(http.StatusOK, gin.H{
		"message": "修改用户信息成功",
	})
}
