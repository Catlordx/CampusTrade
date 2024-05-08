package operation

import (
	"gorm.io/gorm"
)

// ModifyUser
//
//	@Description: 修改用户信息
//	@param	db				数据库DB
//	@param	username		用户名
//	@param	newUsername		新用户名
//	@param	newRealName		新真实姓名
//	@param	newPassword		新密码
//	@param	newPhoneNumber	新手机号
//	@return	bool			修改结果
func ModifyUser(db *gorm.DB, username, newUsername, newRealName, newPassword, newPhoneNumber string) bool {
	user := User(db, username)
	if user == nil {
		return false
	}

	if newRealName != "" {
		ModifyRealName(db, username, newRealName)
	}
	if newPassword != "" {
		ModifyPassword(db, username, newPassword)
	}
	if newPhoneNumber != "" {
		ModifyPhoneNumber(db, username, newPhoneNumber)
	}
	if newUsername != "" {
		ModifyUsername(db, username, newUsername)
	}
	return true
}

// ModifyUsername
//
//	@Description: 修改用户的用户名
//	@param	db			数据库DB
//	@param	username	用户名
//	@param	newUsername	新用户名
//	@return	bool		修改结果
func ModifyUsername(db *gorm.DB, username, newUsername string) bool {
	user := User(db, username)
	if user == nil {
		return false
	}
	db.Model(&user).Update("username", newUsername)
	return true
}

// ModifyRealName
//
//	@Description: 修改用户的真实姓名
//	@param	db			数据库DB
//	@param	username	用户名
//	@param	newRealName	新真实姓名
//	@return	bool		修改结果
func ModifyRealName(db *gorm.DB, username, newRealName string) bool {
	user := User(db, username)
	if user == nil {
		return false
	}
	db.Model(&user).Update("real_name", newRealName)
	return true
}

// ModifyPassword
//
//	@Description: 修改用户的密码
//	@param	db			数据库DB
//	@param	username	用户名
//	@param	newPassword	新密码
//	@return	bool		修改结果
func ModifyPassword(db *gorm.DB, username, newPassword string) bool {
	user := User(db, username)
	if user == nil {
		return false
	}
	db.Model(&user).Update("password", newPassword)
	return true
}

// ModifyPhoneNumber
//
//	@Description: 修改用户的手机号
//	@param	db				数据库DB
//	@param	username		用户名
//	@param	newPhoneNumber	新手机号
//	@return	bool			修改结果
func ModifyPhoneNumber(db *gorm.DB, username, newPhoneNumber string) bool {
	user := User(db, username)
	if user == nil {
		return false
	}
	db.Model(&user).Update("phone_number", newPhoneNumber)
	return true
}
