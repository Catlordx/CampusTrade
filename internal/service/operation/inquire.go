package operation

import (
	"errors"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"gorm.io/gorm"
)

// User
//
//	@Description: 查询用户
//	@param	db			数据库DB
//	@param	username	查询用户名
//	@return	*mysql.User	用户结构体
func User(db *gorm.DB, username string) *mysql.User {
	var user mysql.User
	result := db.First(&user, "username = ?", username)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

// UserRealName
//
//	@Description: 查询用户真实名字
//	@param	db			数据库DB
//	@param	username	查询用户真实名字
//	@return	string		真实名字字符串
func UserRealName(db *gorm.DB, username string) string {
	user := User(db, username)
	if user == nil {
		return ""
	}
	return user.RealName
}

// UserPassword
//
//	@Description: 查询用户密码
//	@param	db			数据库DB
//	@param	username	查询用户名
//	@return	string		密码字符串
func UserPassword(db *gorm.DB, username string) string {
	user := User(db, username)
	if user == nil {
		return ""
	}
	return user.Password
}

// UserPhoneNumber
//
//	@Description: 查询用户手机号
//	@param	db			数据库DB
//	@param	username	查询用户名
//	@return	string		手机号字符串
func UserPhoneNumber(db *gorm.DB, username string) string {
	user := User(db, username)
	if user == nil {
		return ""
	}
	return user.PhoneNumber
}

// UserRoleID
//
//	@Description: 查询用户角色ID
//	@param	db			数据库DB
//	@param	username
//	@return	uint
func UserRoleID(db *gorm.DB, username string) uint {
	user := User(db, username)
	if user == nil {
		return 0
	}
	return user.RoleID
}

// HasPermission
//
//	@Description: 查询用户是否具有权限
//	@param	db				数据库DB
//	@param	username		用户名
//	@param	permissionName	权限字符串
//	@return	bool			查询结果
func HasPermission(db *gorm.DB, username string, permissionID uint) bool {
	user := User(db, username)
	if user == nil {
		return false
	}
	permissions := RolePermissionID(db, user.RoleID)
	for _, permission := range permissions {
		if permission == permissionID {
			return true
		}
	}
	return false
}

/*
// RoleID
//
//	@Description: 查询角色ID
//	@param	db			数据库DB
//	@param	roleName	角色字符串
//	@return	uint		角色ID
func RoleID(db *gorm.DB, roleName string) uint {
	roleName = strings.ToLower(roleName)
	var role mysql.Role
	result := db.Where("role_name = ?", roleName).First(&role)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 0
	}
	return role.ID
}

// RoleName
//
//	@Description: 查询角色名
//	@param	db		数据库DB
//	@param	roleID	角色ID
//	@return	string	角色字符串
func RoleName(db *gorm.DB, roleID uint) string {
	var role mysql.Role
	result := db.First(&role, roleID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ""
	}
	return role.Name
}
*/

// RolePermissionID
//
//	@Description: 查询角色具有的权限
//	@param	db		数据库DB
//	@param	roleID	角色ID
//	@return	[]uint	权限切片
func RolePermissionID(db *gorm.DB, roleID uint) []uint {
	var permissionIDs []uint
	db.Model(&mysql.RolePermission{}).Where("role_id = ?", roleID).Pluck("permission_id", &permissionIDs)
	if len(permissionIDs) == 0 {
		return nil
	}
	return permissionIDs
}

/*
// PermissionID
//
//	@Description: 查询权限ID
//	@param	db				数据库DB
//	@param	permissionName	权限字符串
//	@return	uint			权限ID
func PermissionID(db *gorm.DB, permissionName string) uint {
	var permission mysql.Permission
	result := db.Where("permission_name = ?", permissionName).First(&permission)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 0
	}
	return permission.ID
}

// PermissionName
//
//	@Description: 查询权限名
//	@param	db				数据库DB
//	@param	permissionID	权限ID
//	@return	string			权限字符串
func PermissionName(db *gorm.DB, permissionID uint) string {
	var permission mysql.Permission
	result := db.First(&permission, permissionID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ""
	}
	return permission.Name
}
*/
