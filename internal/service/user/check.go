package user

import (
	"github.com/Catlordx/CampusTrade/internal/core/config"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	_user "github.com/Catlordx/CampusTrade/internal/db/mysql/user"
	"github.com/Catlordx/CampusTrade/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserFromClaims
//
//	@Description: 从中间件获取claims并返回对应的mysql.User结构体
//	@param	c			*gin.Context
//	@return	*mysql.User	用户信息结构体
func GetUserFromClaims(c *gin.Context) *mysql.User {
	claims, _ := c.Get("claims")
	userClaims, ok := claims.(*utils.CustomClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "claims类型错误"})
		return nil
	}

	appContext := c.MustGet("appContext").(*config.AppContext)
	user := _user.GetUserByID(appContext.DB, userClaims.UserID)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "未找到用户信息"})
		return nil
	}

	return user
}

// CheckUserPermission
//
//	@Description: 检查用户是否具有权限，并返回处理结果
//	@param	c			*gin.Context
//	@param	user		用户结构体
//	@param	permission  权限
//	@return	bool		处理结果
func CheckUserPermission(c *gin.Context, user *mysql.User, permission string) bool {
	appContext := c.MustGet("appContext").(*config.AppContext)
	if _user.HasPermission(appContext.DB, user, permission) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有“" + permission + "”权限",
		})
		return false
	}
	return true
}

// CheckModifyUser
//
//	@Description: 检查用户处理结果
//	@param	c				*gin.Context
//	@param	userID			被修改用户ID
//	@param	newUsername		新用户名
//	@param	newRealName		新真实姓名
//	@param	newPassword		新密码
//	@param	newPhoneNumber	新手机号
func CheckModifyUser(c *gin.Context, userID uint, newUsername, newRealName, newPassword, newPhoneNumber string) {
	var err error
	appContext := c.MustGet("appContext").(*config.AppContext)

	//修改用户名，newUsername为空字符串时表示不修改用户名
	if newUsername != "" {
		err = _user.ModifyUsername(appContext.DB, userID, newUsername)
		if err != nil {
			switch err.Error() {
			case "new username is the same as old username":
				c.JSON(http.StatusBadRequest, gin.H{"message": "新用户名与旧用户名相同"})
			default:
				c.JSON(http.StatusBadRequest, gin.H{"message": "未知的错误"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "修改用户名成功"})
		}
	}

	//修改真实姓名，newRealName为空字符串时表示不修改真实姓名
	if newRealName != "" {
		err = _user.ModifyRealName(appContext.DB, userID, newRealName)
		if err != nil {
			switch err.Error() {
			case "new real name is the same as old real name":
				c.JSON(http.StatusBadRequest, gin.H{"message": "新真实姓名与旧真实姓名相同"})
			default:
				c.JSON(http.StatusBadRequest, gin.H{"message": "未知的错误"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "修改真实姓名成功"})
		}
	}

	//修改密码，newPassword为空字符串时表示不修改密码
	if newPassword != "" {
		err = _user.ModifyPassword(appContext.DB, userID, newPassword)
		if err != nil {
			switch err.Error() {
			case "new password is the same as old password":
				c.JSON(http.StatusBadRequest, gin.H{"message": "新密码与旧密码相同"})
			default:
				c.JSON(http.StatusBadRequest, gin.H{"message": "未知的错误"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "修改密码成功"})
		}
	}

	//修改手机号，newPhoneNumber为空字符串时表示不修改手机号
	if newPhoneNumber != "" {
		err = _user.ModifyPhoneNumber(appContext.DB, userID, newPhoneNumber)
		if err != nil {
			switch err.Error() {
			case "new phone number is the same as old phone number":
				c.JSON(http.StatusBadRequest, gin.H{"message": "新手机号与旧手机号相同"})
			default:
				c.JSON(http.StatusBadRequest, gin.H{"message": "未知的错误"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "修改手机号成功"})
		}
	}
}
