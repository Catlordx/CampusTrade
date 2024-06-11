package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ModifyUsername
//
//	@Description: 修改用户的用户名
//	@param	db			数据库DB
//	@param	userID		用户ID
//	@param	newUsername	新用户名
//	@return	error		修改结果
func ModifyUsername(db *gorm.DB, userID uint, newUsername string) error {
	user := GetUserByID(db, userID)
	if user.Username == newUsername {
		return errors.New("new username is the same as old username")
	}
	db.Model(&user).Update("username", newUsername)
	return nil
}

// ModifyRealName
//
//	@Description: 修改用户的真实姓名
//	@param	db			数据库DB
//	@param	userID		用户ID
//	@param	newRealName	新真实姓名
//	@return	error		修改结果
func ModifyRealName(db *gorm.DB, userID uint, newRealName string) error {
	user := GetUserByID(db, userID)
	if newRealName == user.RealName {
		return errors.New("new real name is the same as old real name")
	}
	db.Model(&user).Update("real_name", newRealName)
	return nil
}

// ModifyPassword
//
//	@Description: 修改用户的密码
//	@param	db			数据库DB
//	@param	userID		用户ID
//	@param	newPassword	新密码
//	@return	error		修改结果
func ModifyPassword(db *gorm.DB, userID uint, newPassword string) error {
	user := GetUserByID(db, userID)
	if CheckPassword(user.Password, newPassword) {
		return errors.New("new password is the same as old password")
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	db.Model(&user).Update("password", hashedPassword)
	return nil
}

// ModifyPhoneNumber
//
//	@Description: 修改用户的手机号
//	@param	db				数据库DB
//	@param	userID			用户ID
//	@param	newPhoneNumber	新手机号
//	@return	error			修改结果
func ModifyPhoneNumber(db *gorm.DB, userID uint, newPhoneNumber string) error {
	user := GetUserByID(db, userID)
	if newPhoneNumber == user.PhoneNumber {
		return errors.New("new phone number is the same as old phone number")
	}
	db.Model(&user).Update("phone_number", newPhoneNumber)
	return nil
}
