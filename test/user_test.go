package test

import (
	"fmt"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/role"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/status"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/user"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddUser(t *testing.T) {
	conf := mysql.DbConfig{}
	db, _ := mysql.Connect(&conf)

	_ = db.AutoMigrate(&mysql.User{})

	testUser1Username := "test_user1"
	testUser1RealName := "TestUser1"
	testUser1Password := "123456"
	testUser1PhoneNumber := "12312341234"
	testUser1Role := role.User
	err := user.AddUser(db,
		testUser1Username,
		testUser1RealName,
		testUser1Password,
		testUser1PhoneNumber,
		testUser1Role)
	if err != nil {
		if err.Error() == "user already exists" {
			t.Errorf("user already exists")
		} else {
			t.Errorf("unknown error: %s", err.Error())
		}
	}

	err = user.AddUser(db,
		testUser1Username,
		testUser1RealName,
		testUser1Password,
		testUser1PhoneNumber,
		testUser1Role)
	require.Error(t, err, "user should be added")

	testUser1 := user.GetUserByUsername(db, testUser1Username)
	db.Unscoped().Delete(&testUser1)
}

func TestModifyUser(t *testing.T) {
	conf := mysql.DbConfig{}
	db, _ := mysql.Connect(&conf)

	_ = db.AutoMigrate(&mysql.User{})
	_ = user.AddUser(db,
		"test_user1",
		"TestUser1",
		"123456",
		"12312341234",
		role.User)

	testUser1 := user.GetUserByUsername(db, "test_user1")
	fmt.Printf("%v\t%v\t%v\t%v\n", testUser1.Username, testUser1.RealName, testUser1.PhoneNumber, testUser1.Password)
	testUser1ID := testUser1.ID

	err := user.ModifyUsername(db, testUser1ID, "test_user1")
	require.Error(t, err, "new username is the same as old username, user should not be modified")
	err = user.ModifyUsername(db, testUser1ID, "test_user1.0")
	require.NoError(t, err, "user should be modified")
	testUser1 = user.GetUserByID(db, testUser1ID)
	fmt.Printf("%v\t%v\t%v\t%v\n", testUser1.Username, testUser1.RealName, testUser1.PhoneNumber, testUser1.Password)

	err = user.ModifyRealName(db, testUser1ID, "TestUser1")
	require.Error(t, err, "new real name is the same as old real name, user should not be modified")
	err = user.ModifyRealName(db, testUser1ID, "TestUser1.0")
	require.NoError(t, err, "user should be modified")
	testUser1 = user.GetUserByID(db, testUser1ID)
	fmt.Printf("%v\t%v\t%v\t%v\n", testUser1.Username, testUser1.RealName, testUser1.PhoneNumber, testUser1.Password)

	err = user.ModifyPhoneNumber(db, testUser1ID, "12312341234")
	require.Error(t, err, "new phone number is the same as old phone number, user should not be modified")
	err = user.ModifyPhoneNumber(db, testUser1ID, "9879869876")
	require.NoError(t, err, "user should be modified")
	testUser1 = user.GetUserByID(db, testUser1ID)
	fmt.Printf("%v\t%v\t%v\t%v\n", testUser1.Username, testUser1.RealName, testUser1.PhoneNumber, testUser1.Password)

	err = user.ModifyPassword(db, testUser1ID, "123456")
	require.Error(t, err, "new password is the same as old password, user should not be modified")
	err = user.ModifyPassword(db, testUser1ID, "asdfgh")
	require.NoError(t, err, "user should be modified")
	testUser1 = user.GetUserByID(db, testUser1ID)
	fmt.Printf("%v\t%v\t%v\t%v\n", testUser1.Username, testUser1.RealName, testUser1.PhoneNumber, testUser1.Password)
}

func TestGetUserInfoList(t *testing.T) {
	conf := mysql.DbConfig{}
	db, _ := mysql.Connect(&conf)
	_ = db.AutoMigrate(&mysql.User{})

	testUser := []mysql.User{
		{Username: "test1", RealName: "N", Role: role.User, Status: status.Offline, Password: []byte(""), PhoneNumber: ""},
		{Username: "test2", RealName: "G", Role: role.User, Status: status.Online, Password: []byte(""), PhoneNumber: ""},
		{Username: "test3", RealName: "Y", Role: role.Admin, Status: status.Online, Password: []byte(""), PhoneNumber: ""},
		{Username: "test4", RealName: "A", Role: role.User, Status: status.Offline, Password: []byte(""), PhoneNumber: ""},
		{Username: "test5", RealName: "O", Role: role.Admin, Status: status.Offline, Password: []byte(""), PhoneNumber: ""},
	}
	db.Model(&mysql.User{}).Create(testUser)
	defer func() {
		for _, i := range testUser {
			_user := user.GetUserByUsername(db, i.Username)
			db.Unscoped().Delete(&_user)
		}
	}()

	ListInfo1 := user.ListUserInfo{
		Class: "username",
		Order: "DESC",
		Page:  2,
		Count: 2,
	}
	userInfos1 := user.GetUserList(db, ListInfo1)
	require.NotEmpty(t, userInfos1, "userInfos1 should not be empty")
	fmt.Println("Username, DESC, Page2, 2 per Page")
	for _, userInfo := range userInfos1 {
		fmt.Printf("Username:%v\n", userInfo.Username)
	}
	fmt.Println()

	ListInfo2 := user.ListUserInfo{
		Class: "real_name",
		Order: "ASC",
		Page:  1,
		Count: 10,
	}
	userInfos2 := user.GetUserList(db, ListInfo2)
	require.NotEmpty(t, userInfos2, "userInfos2 should not be empty")
	fmt.Println("RealName, ASC, Page1, 10 per Page")
	for _, userInfo := range userInfos2 {
		fmt.Printf("Username:%v\tRealName:%v\n", userInfo.Username, userInfo.RealName)
	}
	fmt.Println()

	ListInfo3 := user.ListUserInfo{
		Class: "role",
		Order: "",
		Page:  1,
		Count: 10,
	}
	userInfos3 := user.GetUserList(db, ListInfo3)
	require.NotEmpty(t, userInfos3, "userInfos3 should not be empty")
	fmt.Println("Role, Default, Page1, 10 per Page")
	for _, userInfo := range userInfos3 {
		fmt.Printf("Username:%v\tRole:%v\n", userInfo.Username, userInfo.Role)
	}
	fmt.Println()

	ListInfo4 := user.ListUserInfo{
		Class: "status",
		Order: "",
		Page:  1,
		Count: 10,
	}
	userInfos4 := user.GetUserList(db, ListInfo4)
	require.NotEmpty(t, userInfos4, "userInfos4 should not be empty")
	fmt.Println("Status, Default, Page1, 10 per Page")
	for _, userInfo := range userInfos4 {
		fmt.Printf("Username:%v\tStatus:%v\n", userInfo.Username, userInfo.Status)
	}
	fmt.Println()

	ListInfo5 := user.ListUserInfo{
		Class: "",
		Order: "",
		Page:  1,
		Count: 10,
	}
	userInfos5 := user.GetUserList(db, ListInfo5)
	require.NotEmpty(t, userInfos5, "userInfos5 should not be empty")
	fmt.Println("Default, Default, Page1, 10 per Page")
	for _, userInfo := range userInfos5 {
		fmt.Printf("Username:%v\n", userInfo.Username)
	}
	fmt.Println()
}
