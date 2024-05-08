package user

import (
	"errors"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"gorm.io/gorm"
)

// GetUserByID
//
//	@Description: 根据ID查询用户
//	@param	db	数据库DB
//	@param	ID	查询用户ID
//	@return	*mysql.GetUserByUsername	用户结构体
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
//	@return	*mysql.GetUserByUsername	用户结构体
func GetUserByUsername(db *gorm.DB, username string) *mysql.User {
	var user mysql.User
	result := db.First(&user, "username = ?", username)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

// HasPermission
//
//	@Description: 查询用户是否具有权限
//	@param	db			数据库DB
//	@param	username	用户名
//	@param	permission	权限字符串
//	@return	bool		查询结果
func HasPermission(db *gorm.DB, username string, permission string) bool {
	user := GetUserByUsername(db, username)
	if user == nil {
		return false
	}
	permissions := RolePermission(db, user.Role)
	for _, p := range permissions {
		if p == permission {
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
	var permissions []string
	db.Model(&mysql.RolePermission{}).Where("role = ?", role).Pluck("permission", &permissions)
	if len(permissions) == 0 {
		return nil
	}
	return permissions
}
