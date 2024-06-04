package user

import (
	"errors"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"golang.org/x/crypto/bcrypt"
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
func HasPermission(db *gorm.DB, role string, permission string) bool {
	permissions := RolePermission(db, role)
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
