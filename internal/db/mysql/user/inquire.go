package user

import (
	"errors"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/role"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/status"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

type ListUserInfo struct {
	Class string `json:"class" form:"class"`
	Order string `json:"order" form:"order"`
	Page  int    `json:"page" form:"page"`
	Count int    `json:"count" form:"count"`
}

type InquireUserInfo struct {
	Username    string
	RealName    string
	PhoneNumber string
	Role        string
	Status      string
}

// GetUserByID
//
//	@Description: 根据ID查询用户
//	@param	db			数据库DB
//	@param	ID			查询用户ID
//	@return	*mysql.USer	用户结构体
func GetUserByID(db *gorm.DB, ID uint) *mysql.User {
	var user mysql.User
	result := db.First(&user, ID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

// GetUserByUsername
//
//	@Description: 根据用户名查询用户
//	@param	db			数据库DB
//	@param	username	查询用户名
//	@return	*mysql.User	用户结构体
func GetUserByUsername(db *gorm.DB, username string) *mysql.User {
	var user mysql.User
	result := db.First(&user, "username = ?", username)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

// CheckPassword
//
//	@Description: 检验输入密码是否与用户密码相同
//	@param	userPassword	加密后的用户密码
//	@param	password		输入密码
//	@return	bool			判断结果
func CheckPassword(userPassword []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(userPassword, []byte(password))
	if err != nil {
		return false
	}
	return true
}

// HasPermission
//
//	@Description: 查询用户是否具有权限
//	@param	db			数据库DB
//	@param	role		角色名
//	@param	permission	权限字符串
//	@return	bool		查询结果
func HasPermission(db *gorm.DB, user *mysql.User, userPermission string) bool {
	permissions := RolePermission(db, user.Role)
	for _, permission := range permissions {
		if permission == userPermission {
			return true
		}
	}
	return false
}

// RolePermission
//
//	@Description: 查询角色具有的权限
//	@param	db			数据库DB
//	@param	role		角色ID
//	@return	[]string	权限切片
func RolePermission(db *gorm.DB, role string) []string {
	if role == "" {
		return []string{}
	}
	var permissions []string
	db.Model(&mysql.RolePermission{}).Where("role = ?", role).Pluck("permission", &permissions)
	if len(permissions) == 0 {
		return nil
	}
	return permissions
}

// GetUserList
//
//	@Description: 获取用户信息列表，信息包括用户名，真实姓名，手机号，用户角色，用户状态
//	@param	db					数据库DB
//	@param	listInfo			用户列表限制要求
//	@return	[]InquireUserInfo	用户信息列表
func GetUserList(db *gorm.DB, listInfo ListUserInfo) []InquireUserInfo {
	var userInfoList []InquireUserInfo

	var order string
	switch listInfo.Class {
	case "":
		order = "id"
	//按role排序：升序为admin-user，降序为user-admin
	case "role":
		order = "CASE " +
			"WHEN role = '" + role.Admin + "' THEN 1 " +
			"WHEN role = '" + role.User + "' THEN 2 " +
			"END"
	//按status排序：升序为online-offline，降序为offline-online
	case "status":
		order = "CASE " +
			"WHEN status = '" + status.Online + "' THEN 1 " +
			"WHEN status = '" + status.Offline + "' THEN 2 " +
			"END"
	default:
		order = listInfo.Class
	}

	if listInfo.Order == "" || strings.ToUpper(listInfo.Order) == "ASC" {
		order += " ASC"
	} else {
		order += " DESC"
	}

	offset := (listInfo.Page - 1) * listInfo.Count
	db.Model(&mysql.User{}).
		Select("username",
			"real_name",
			"phone_number",
			"role",
			"status").
		Order(order).
		Offset(offset).
		Limit(listInfo.Count).
		Find(&userInfoList)

	return userInfoList
}
