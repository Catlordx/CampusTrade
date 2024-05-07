package user

import (
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
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
		"role":         permission.String(user.RoleID),
	})
}

// ModifyUsername
//
//	@Description: 用户修改用户名
//	@param	c	gin.Context
func ModifyUsername(c *gin.Context) {
	username := c.PostForm("username")
	newUsername := c.PostForm("new_username")

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
			"message": "未找到用户信息",
		})
		return
	}
	if operation.HasPermission(db, username, permission.ModifyUsername) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改自己用户名的权限",
		})
		return
	}
	if _, err := operation.ModifyUsername(db, username, newUsername); err != nil {
		switch err.Error() {
		case "modified username is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入被修改用户的用户名",
			})
		case "user is not exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未找到用户信息",
			})
		case "new username is the same as old username":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "新用户名不应与旧用户名相同",
			})
		case "username is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入新用户名",
			})
		case "username is too short":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "新用户名过短",
			})
		case "username is invalid":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "新用户名不应包含特殊字符",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未知的错误",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改用户名成功",
	})
}

// ModifyRealName
//
//	@Description: 用户修改真实名字
//	@param	c	gin.Context
func ModifyRealName(c *gin.Context) {
	username := c.PostForm("username")
	newRealName := c.PostForm("new_real_name")

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
			"message": "未找到用户信息",
		})
		return
	}
	if operation.HasPermission(db, username, permission.ModifyRealName) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改自己真实名字的权限",
		})
		return
	}
	if _, err := operation.ModifyRealName(db, username, newRealName); err != nil {
		switch err.Error() {
		case "modified username is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入被修改用户的用户名",
			})
		case "user is not exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未找到用户信息",
			})
		case "new real name is the same as old real name":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "新真实名字不应与旧真实名字相同",
			})
		case "real name is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入新真实姓名",
			})
		case "real name is invalid":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "新真实名字不应包含特殊字符",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未知的错误",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改真实名字成功",
	})
}

// ModifyPassword
//
//	@Description: 用户修改密码
//	@param	c	gin.Context
func ModifyPassword(c *gin.Context) {
	username := c.PostForm("username")
	newPassword := c.PostForm("new_password")

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
			"message": "未找到用户信息",
		})
		return
	}
	if operation.HasPermission(db, username, permission.ModifyPassword) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改自己密码的权限",
		})
		return
	}
	if _, err := operation.ModifyPassword(db, username, newPassword); err != nil {
		switch err.Error() {
		case "modified username is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入被修改用户的用户名",
			})
		case "user is not exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未找到用户信息",
			})
		case "new password is the same as old password":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "新密码不应与旧密码相同",
			})
		case "password is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入新密码",
			})
		case "password is too short":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "新密码过短",
			})
		case "password is invalid":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "新密码不应包含特殊字符",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未知的错误",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改密码成功",
	})
}

// ModifyPhoneNumber
//
//	@Description: 用户修改手机号
//	@param	c	gin.Context
func ModifyPhoneNumber(c *gin.Context) {
	username := c.PostForm("username")
	newPhoneNumber := c.PostForm("new_phone_number")

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
			"message": "未找到用户信息",
		})
		return
	}
	if operation.HasPermission(db, username, permission.ModifyPhoneNumber) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改自己手机号的权限",
		})
		return
	}
	if _, err := operation.ModifyPhoneNumber(db, username, newPhoneNumber); err != nil {
		switch err.Error() {
		case "modified username is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入被修改用户的用户名",
			})
		case "user is not exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未找到用户信息",
			})
		case "new phone number is the same as old phone number":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "新手机号不应与旧手机号相同",
			})
		case "phone number is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入新手机号",
			})
		case "phone number is invalid":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请输入正确格式的手机号",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未知的错误",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改手机号成功",
	})
}

// ModifyRole
//
//	@Description: 用户修改角色
//	@param	c	gin.Context
func ModifyRole(c *gin.Context) {
	username := c.PostForm("username")
	newRole := c.PostForm("new_role")

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
			"message": "未找到用户信息",
		})
		return
	}
	if operation.HasPermission(db, username, permission.ModifyRole) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改自己角色的权限",
		})
		return
	}
	if _, err := operation.ModifyRole(db, username, newRole); err != nil {
		switch err.Error() {
		case "modified username is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入被修改用户的用户名",
			})
		case "user is not exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未找到用户信息",
			})
		case "new role is the same as old role":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "新角色不应与旧角色相同",
			})
		case "role is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入新角色",
			})
		case "role is not exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "不存在该角色",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未知的错误",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改角色成功",
	})
}
