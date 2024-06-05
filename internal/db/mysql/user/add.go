package user

import (
	"errors"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AddUser
//
//	@Description: 添加用户
//	@param	db			数据库DB
//	@param	username	用户名
//	@param	realName	真实名字
//	@param	password	密码
//	@param	phoneNumber	手机号
//	@param	role    	角色
//	@return	error		添加结果
func AddUser(db *gorm.DB, username, realName, password, phoneNumber, role string) error {
	if user := GetUserByUsername(db, username); user != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := mysql.User{
		Username:    username,
		RealName:    realName,
		Password:    hashedPassword,
		PhoneNumber: phoneNumber,
		Role:        role,
	}
	db.Create(&user)
	return nil
}
