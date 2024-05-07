package operation

import (
	"errors"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/role"
	"github.com/Catlordx/CampusTrade/internal/service/tool"
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
//	@param	roleName	角色
//	@return	bool		添加结果
//	@return	error		失败原因
func AddUser(db *gorm.DB, username, realName, password, phoneNumber, roleName string) (bool, error) {
	if user := User(db, username); user != nil {
		return false, errors.New("user already exists")
	}
	if _, err := tool.CheckUsername(username); err != nil {
		return false, err
	}
	if _, err := tool.CheckRealName(realName); err != nil {
		return false, err
	}
	if _, err := tool.CheckPassword(password); err != nil {
		return false, err
	}
	if _, err := tool.CheckPhoneNumber(phoneNumber); err != nil {
		return false, err
	}
	if _, err := tool.CheckRole(roleName); err != nil {
		return false, err
	}
	user := mysql.User{
		Username:    username,
		RealName:    realName,
		Password:    password,
		PhoneNumber: phoneNumber,
		RoleID:      role.ID(roleName),
	}
	db.Create(&user)
	return true, nil
}
