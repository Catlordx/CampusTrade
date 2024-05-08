package operation

import (
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
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
//	@return	bool		添加结果
func AddUser(db *gorm.DB, username, realName, password, phoneNumber, role string) bool {
	if user := User(db, username); user != nil {
		return false
	}
	user := mysql.User{
		Username:    username,
		RealName:    realName,
		Password:    password,
		PhoneNumber: phoneNumber,
		Role:        role,
	}
	db.Create(&user)
	return true
}
