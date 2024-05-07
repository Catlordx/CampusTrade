package tool

import (
	"errors"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/role"
	"regexp"
	"strings"
)

// CheckUsername
//
//	@Description: 检查名字是否合规
//	@param	username	名字字符串
//	@return	bool		是否合规
//	@return	error		不合格原因
func CheckUsername(username string) (bool, error) {
	if username == "" {
		return false, errors.New("username is empty")
	}
	//判断用户名是否过短
	if len(username) < 2 {
		return false, errors.New("username is too short")
	}

	//判断用户名是否包含特殊字符
	match, err := regexp.MatchString("^[a-zA-Z0-9\\x{4e00}-\\x{9fa5}._·]+$", username)
	if err != nil {
		return false, err
	}
	if match == false {
		return false, errors.New("username is invalid")
	}

	return true, nil
}

// CheckRealName
//
//	@Description: 检查名字是否合规
//	@param	realName	名字字符串
//	@return	bool		是否合规
//	@return	error		不合格原因
func CheckRealName(realName string) (bool, error) {
	if realName == "" {
		return false, errors.New("real name is empty")
	}

	//判断用户名是否包含特殊字符
	match, err := regexp.MatchString("^[a-zA-Z0-9\\x{4e00}-\\x{9fa5}._·]+$", realName)
	if err != nil {
		return false, err
	}
	if match == false {
		return false, errors.New("real name is invalid")
	}

	return true, nil
}

// CheckPassword
//
//	@Description: 检查密码是否合规
//	@param	password	密码字符串
//	@return	bool		是否合规
//	@return	error		不合规原因
func CheckPassword(password string) (bool, error) {
	if password == "" {
		return false, errors.New("password is empty")
	}
	//判断密码是否过短
	if len(password) < 6 {
		return false, errors.New("password is too short")
	}

	//判断密码是否包含特殊字符
	match, err := regexp.MatchString("^[a-zA-Z0-9]+$", password)
	if err != nil {
		return false, err
	}
	if match == false {
		return false, errors.New("password is invalid")
	}

	//密码合规
	return true, nil
}

// CheckPhoneNumber
//
//	@Description: 检查手机号是否合规
//	@param	phoneNumber	手机号字符串
//	@return	bool		是否合规
//	@return	error		不合规原因
func CheckPhoneNumber(phoneNumber string) (bool, error) {
	if phoneNumber == "" {
		return false, errors.New("phone number is empty")
	}
	//判断手机号是否包含特殊字符
	match, err := regexp.MatchString("^1\\d{10}$", phoneNumber)
	if err != nil {
		return false, err
	}
	if match == false {
		return false, errors.New("phone number is invalid")
	}

	//手机号合规
	return true, nil
}

// CheckRole
//
//	@Description: 检查角色是否合规
//	@param	role	角色字符串
//	@return	string	返回全小写的角色字符串
//	@return	error	不合规原因
func CheckRole(roleName string) (string, error) {
	if roleName == "" {
		return "", errors.New("role is empty")
	}
	//判断角色是否合规
	roleName = strings.ToLower(roleName)
	if roleID := role.ID(roleName); roleID == 0 {
		return "", errors.New("role is not exist")
	}
	return roleName, nil
}
