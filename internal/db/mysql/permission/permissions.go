package permission

//goland:noinspection GoCommentStart
const (
	//ID: 0 无权限
	//ID: 1-9 创造信息权限
	AddUser uint = 1
	//ID: 10-19 个人信息权限
	InquireInfo       uint = 10
	ModifyUsername    uint = 11
	ModifyRealName    uint = 12
	ModifyPassword    uint = 13
	ModifyPhoneNumber uint = 14
	ModifyRole        uint = 15
	//ID: 20-29 用户信息权限
	InquireAnyoneInfo       uint = 20
	ModifyAnyoneUsername    uint = 21
	ModifyAnyoneRealName    uint = 22
	ModifyAnyonePassword    uint = 23
	ModifyAnyonePhoneNumber uint = 24
	ModifyAnyoneRole        uint = 25
)

func String(permissionID uint) string {
	switch permissionID {
	case AddUser:
		return "add_user"

	case InquireInfo:
		return "inquire_info"
	case ModifyUsername:
		return "modify_username"
	case ModifyRealName:
		return "modify_real_name"
	case ModifyPassword:
		return "modify_password"
	case ModifyPhoneNumber:
		return "modify_phone_number"
	case ModifyRole:
		return "modify_role"

	case InquireAnyoneInfo:
		return "inquire_anyone_info"
	case ModifyAnyoneUsername:
		return "modify_anyone_username"
	case ModifyAnyoneRealName:
		return "modify_anyone_real_name"
	case ModifyAnyonePassword:
		return "modify_anyone_password"
	case ModifyAnyonePhoneNumber:
		return "modify_anyone_phone_number"
	case ModifyAnyoneRole:
		return "modify_anyone_role"
	}
	return ""
}

func ID(permissionName string) uint {
	switch permissionName {
	case "add_user":
		return AddUser

	case "inquire_info":
		return InquireInfo
	case "modify_username":
		return ModifyUsername
	case "modify_real_name":
		return ModifyRealName
	case "modify_password":
		return ModifyPassword
	case "modify_phone_number":
		return ModifyPhoneNumber
	case "modify_role":
		return ModifyRole

	case "inquire_anyone_info":
		return InquireAnyoneInfo
	case "modify_anyone_username":
		return ModifyAnyoneUsername
	case "modify_anyone_real_name":
		return ModifyAnyoneRealName
	case "modify_anyone_password":
		return ModifyAnyonePassword
	case "modify_anyone_phone_number":
		return ModifyAnyonePhoneNumber
	case "modify_anyone_role":
		return ModifyAnyoneRole
	}
	return 0
}
