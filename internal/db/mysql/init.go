package mysql

import (
	"github.com/Catlordx/CampusTrade/internal/db/mysql/permission"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/role"
	"gorm.io/gorm"
)

// InitRolePermission
//
//	@Description:	初始化角色权限表
//					用户角色权限包括：
//						查询自己信息，信息包括真实名字、手机号和角色；
//						修改自己用户名；
//						修改自己真实姓名；
//						修改自己密码；
//						修改自己手机号。
//					管理员角色权限包括：
//						查询自己信息，信息包括真实名字、手机号和角色；
//						修改自己用户名；
//						修改自己真实姓名；
//						修改自己密码；
//						修改自己手机号；
//						修改自己角色；
//						查询用户信息，信息包括手机号和角色；
//						修改用户用户名；
//						修改用户真实姓名；
//						修改用户密码；
//						修改用户手机号；
//						修改用户角色。
//	@param	db
func InitRolePermission(db *gorm.DB) {
	db.AutoMigrate(&User{} /*, &Role{}, &Permission{}*/, &RolePermission{})
	//角色权限关联表
	rolePermissions := []RolePermission{
		//ID: 100-199 user角色拥有的权限
		{ID: 110, RoleID: role.User, PermissionID: permission.InquireInfo},
		{ID: 111, RoleID: role.User, PermissionID: permission.ModifyUsername},
		{ID: 112, RoleID: role.User, PermissionID: permission.ModifyRealName},
		{ID: 113, RoleID: role.User, PermissionID: permission.ModifyPassword},
		{ID: 114, RoleID: role.User, PermissionID: permission.ModifyPhoneNumber},
		//ID: 200-299 admin角色拥有的权限
		{ID: 201, RoleID: role.Admin, PermissionID: permission.AddUser},
		{ID: 210, RoleID: role.Admin, PermissionID: permission.InquireInfo},
		{ID: 211, RoleID: role.Admin, PermissionID: permission.ModifyUsername},
		{ID: 212, RoleID: role.Admin, PermissionID: permission.ModifyRealName},
		{ID: 213, RoleID: role.Admin, PermissionID: permission.ModifyPassword},
		{ID: 214, RoleID: role.Admin, PermissionID: permission.ModifyPhoneNumber},
		{ID: 215, RoleID: role.Admin, PermissionID: permission.ModifyRole},
		{ID: 220, RoleID: role.Admin, PermissionID: permission.InquireAnyoneInfo},
		{ID: 221, RoleID: role.Admin, PermissionID: permission.ModifyAnyoneUsername},
		{ID: 222, RoleID: role.Admin, PermissionID: permission.ModifyAnyoneRealName},
		{ID: 223, RoleID: role.Admin, PermissionID: permission.ModifyAnyonePassword},
		{ID: 224, RoleID: role.Admin, PermissionID: permission.ModifyAnyonePhoneNumber},
		{ID: 225, RoleID: role.Admin, PermissionID: permission.ModifyAnyoneRole},
	}
	for _, rolePermission := range rolePermissions {
		db.FirstOrCreate(&rolePermission, RolePermission{ID: rolePermission.ID})
	}
}
