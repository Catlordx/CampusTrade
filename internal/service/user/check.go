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
//	@param	username	用户名
//	@param	permission  权限
//	@return	bool		处理结果
func CheckUserPermission(c *gin.Context, role string, permission string) bool {
	appContext := c.MustGet("appContext").(*config.AppContext)
	if _user.HasPermission(appContext.DB, role, permission) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有“" + permission + "”权限",
		})
		return false
	}
	return true
}
