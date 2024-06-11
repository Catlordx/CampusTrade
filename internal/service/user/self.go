package user

import (
	"github.com/Catlordx/CampusTrade/internal/core/config"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/commodity"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/permission"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// InquireInfo
//
//	@Description: 查询用户自己的信息，返回信息包括用户的真实名字、手机号和用户角色
//	@param	c	*gin.Context
func InquireInfo(c *gin.Context) {
	user := GetUserFromClaims(c)
	if user == nil {
		return
	}

	if CheckUserPermission(c, user, permission.InquireInfo) == false {
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
	user := GetUserFromClaims(c)
	if user == nil {
		return
	}

	if CheckUserPermission(c, user, permission.ModifyInfo) == false {
		return
	}

	newUsername := c.PostForm("new_username")
	newRealName := c.PostForm("new_real_name")
	newPassword := c.PostForm("new_password")
	newPhoneNumber := c.PostForm("new_phone_number")

	CheckModifyUser(c, user.ID, newUsername, newRealName, newPassword, newPhoneNumber)
}

// AddFavorite
//
//	@Description: 添加收藏商品
//	@param	c
func AddFavorite(c *gin.Context) {
	user := GetUserFromClaims(c)
	if user == nil {
		return
	}

	if CheckUserPermission(c, user, permission.OperateFavorite) == false {
		return
	}

	favoriteID, _ := strconv.Atoi(c.PostForm("favorite_id"))

	appContext := c.MustGet("appContext").(*config.AppContext)
	err := commodity.AddFavorite(appContext.DB, user.ID, uint(favoriteID))
	if err != nil {
		switch err.Error() {
		case "commodity not found":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "商品不存在",
			})
		case "commodity already be in favorite":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "商品已经被收藏",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "商品收藏成功",
	})
}

// RemoveFavorite
//
//	@Description: 移除收藏商品
//	@param	c	*gin.Context
func RemoveFavorite(c *gin.Context) {
	user := GetUserFromClaims(c)
	if user == nil {
		return
	}

	if CheckUserPermission(c, user, permission.OperateFavorite) == false {
		return
	}

	appContext := c.MustGet("appContext").(*config.AppContext)

	favoriteID, _ := strconv.Atoi(c.PostForm("favorite_id"))
	result := commodity.RemoveFavorite(appContext.DB, user.ID, uint(favoriteID))
	if result == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "商品未收藏",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "移除收藏成功",
	})
}

// FavoriteList
//
//	@Description: 获取收藏商品列表
//	@param	c	c.Context
func FavoriteList(c *gin.Context) {
	user := GetUserFromClaims(c)
	if user == nil {
		return
	}

	if CheckUserPermission(c, user, permission.OperateFavorite) == false {
		return
	}

	sort := c.PostForm("count")
	reverseStr := c.PostForm("reverse")
	page, _ := strconv.Atoi(c.PostForm("page"))
	count, _ := strconv.Atoi(c.PostForm("count"))

	var reverse bool
	if strings.ToLower(reverseStr) == "asc" || strings.ToLower(reverseStr) == "" { //默认升序
		reverse = false
	} else if strings.ToLower(reverseStr) == "desc" {
		reverse = true
	}

	appContext := c.MustGet("appContext").(*config.AppContext)
	commodityList :=
		commodity.GetFavorites(
			appContext.DB,
			user.ID,
			sort,
			reverse,
			page,
			count)

	c.JSON(http.StatusOK, commodityList)
}
