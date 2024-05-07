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
	if operation.HasPermission(db, username, permission.InquireAnyoneInfo) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有查询其他用户信息的权限",
		})
		return
	}

	if inquiredUsername == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未输入要查询用户的用户名",
		})
		return
	}
	inquiredUser := operation.User(db, inquiredUsername)
	if inquiredUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "查询的用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"massage":      "查询用户信息成功",
		"phone_number": user.PhoneNumber,
		"role":         permission.String(user.RoleID),
	})
}

// ModifyAnyoneUsername
//
//	@Description: 修改其他用户的用户名
//	@param	c	gin.Context
func ModifyAnyoneUsername(c *gin.Context) {
	username := c.PostForm("username")
	modifiedUsername := c.PostForm("modified_username")
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
	if operation.HasPermission(db, username, permission.ModifyAnyoneUsername) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改其他用户的用户名的权限",
		})
		return
	}

	if _, err := operation.ModifyUsername(db, modifiedUsername, newUsername); err != nil {
		switch err.Error() {
		case "modified username is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入被修改用户的用户名",
			})
		case "user is not exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "要修改的用户不存在",
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

// ModifyAnyoneRealName
//
//	@Description: 修改其他用户的真实名字
//	@param	c	gin.Context
func ModifyAnyoneRealName(c *gin.Context) {
	username := c.PostForm("username")
	modifiedUsername := c.PostForm("modified_username")
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
	if operation.HasPermission(db, username, permission.ModifyAnyoneRealName) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改其他用户的真实名字的权限",
		})
		return
	}

	if _, err := operation.ModifyRealName(db, modifiedUsername, newRealName); err != nil {
		switch err.Error() {
		case "modified username is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入被修改用户的用户名",
			})
		case "user is not exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "要修改的用户不存在",
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

// ModifyAnyonePassword
//
//	@Description: 修改其他用户的密码
//	@param	c	gin.Context
func ModifyAnyonePassword(c *gin.Context) {
	username := c.PostForm("username")
	modifiedUsername := c.PostForm("modified_username")
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
	if operation.HasPermission(db, username, permission.ModifyAnyonePassword) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改其他用户的密码的权限",
		})
		return
	}

	if _, err := operation.ModifyPassword(db, modifiedUsername, newPassword); err != nil {
		switch err.Error() {
		case "modified username is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入被修改用户的用户名",
			})
		case "user is not exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "要修改的用户不存在",
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

// ModifyAnyonePhoneNumber
//
//	@Description: 修改其他用户的手机号
//	@param	c	gin.Context
func ModifyAnyonePhoneNumber(c *gin.Context) {
	username := c.PostForm("username")
	modifiedUsername := c.PostForm("modified_username")
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
	if operation.HasPermission(db, username, permission.ModifyAnyonePhoneNumber) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改其他用户的手机号的权限",
		})
		return
	}

	if _, err := operation.ModifyPhoneNumber(db, modifiedUsername, newPhoneNumber); err != nil {
		switch err.Error() {
		case "modified username is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入被修改用户的用户名",
			})
		case "user is not exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "要修改的用户不存在",
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

// ModifyAnyoneRole
//
//	@Description: 修改其他用户的角色
//	@param	c	gin.Context
func ModifyAnyoneRole(c *gin.Context) {
	username := c.PostForm("username")
	modifiedUsername := c.PostForm("modified_username")
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
	if operation.HasPermission(db, username, permission.ModifyAnyoneRole) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不具有修改其他用户的角色的权限",
		})
		return
	}

	if _, err := operation.ModifyRole(db, modifiedUsername, newRole); err != nil {
		switch err.Error() {
		case "modified username is empty":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "未输入被修改用户的用户名",
			})
		case "user is not exist":
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "要修改的用户不存在",
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
