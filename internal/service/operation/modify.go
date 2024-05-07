package operation

import (
	"errors"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/role"
	"github.com/Catlordx/CampusTrade/internal/service/tool"
	"gorm.io/gorm"
)

// ModifyUsername
//
//	@Description: 修改用户的用户名
//	@param	db			数据库DB
//	@param	username	用户名
//	@param	newUsername	新用户名
//	@return	bool		修改结果
//	@return	error		修改失败原因
func ModifyUsername(db *gorm.DB, username, newUsername string) (bool, error) {
	if username == "" {
		return false, errors.New("modified username is empty")
	}
	user := User(db, username)
	if user == nil {
		return false, errors.New("user is not exist")
	}
	if newUsername == user.Username {
		return false, errors.New("new username is the same as old username")
	}
	if _, err := tool.CheckUsername(newUsername); err != nil {
		return false, err
	}
	db.Model(&user).Update("username", newUsername)
	return true, nil
}

// ModifyRealName
//
//	@Description: 修改用户的真实姓名
//	@param	db			数据库DB
//	@param	username	用户名
//	@param	newRealName	新真实姓名
//	@return	bool		修改结果
//	@return	error		修改失败原因
func ModifyRealName(db *gorm.DB, username, newRealName string) (bool, error) {
	if username == "" {
		return false, errors.New("modified username is empty")
	}
	user := User(db, username)
	if user == nil {
		return false, errors.New("user is not exist")
	}
	if newRealName == user.RealName {
		return false, errors.New("new real name is the same as old real name")
	}
	if _, err := tool.CheckRealName(newRealName); err != nil {
		return false, err
	}
	db.Model(&user).Update("real_name", newRealName)
	return true, nil
}

// ModifyPassword
//
//	@Description: 修改用户的密码
//	@param	db			数据库DB
//	@param	username	用户名
//	@param	newPassword	新密码
//	@return	bool		修改结果
//	@return	error		修改失败原因
func ModifyPassword(db *gorm.DB, username, newPassword string) (bool, error) {
	if username == "" {
		return false, errors.New("modified username is empty")
	}
	user := User(db, username)
	if user == nil {
		return false, errors.New("user is not exist")
	}
	if newPassword == user.Password {
		return false, errors.New("new password is the same as old password")
	}
	if _, err := tool.CheckPassword(newPassword); err != nil {
		return false, err
	}
	db.Model(&user).Update("password", newPassword)
	return true, nil
}

// ModifyPhoneNumber
//
//	@Description: 修改用户的手机号
//	@param	db				数据库DB
//	@param	username		用户名
//	@param	newPhoneNumber	新手机号
//	@return	bool			修改结果
//	@return	error			修改失败原因
func ModifyPhoneNumber(db *gorm.DB, username, newPhoneNumber string) (bool, error) {
	if username == "" {
		return false, errors.New("modified username is empty")
	}
	user := User(db, username)
	if user == nil {
		return false, errors.New("user is not exist")
	}
	if newPhoneNumber == user.PhoneNumber {
		return false, errors.New("new phone number is the same as old phone number")
	}
	if _, err := tool.CheckPhoneNumber(newPhoneNumber); err != nil {
		return false, err
	}
	db.Model(&user).Update("phone_number", newPhoneNumber)
	return true, nil
}

// ModifyRole
//
//	@Description: 修改用户的角色
//	@param	db			数据库DB
//	@param	username	用户名
//	@param	newRole		新角色
//	@return	bool		修改结果
//	@return	error		修改失败原因
func ModifyRole(db *gorm.DB, username, newRole string) (bool, error) {
	if username == "" {
		return false, errors.New("modified username is empty")
	}
	user := User(db, username)
	if user == nil {
		return false, errors.New("user is not exist")
	}
	if role.ID(newRole) == user.RoleID {
		return false, errors.New("new role is the same as old role")
	}
	if _, err := tool.CheckRole(newRole); err != nil {
		return false, err
	}
	newRoleID := role.ID(newRole)
	db.Model(&user).Update("role_id", newRoleID)
	return true, nil
}
