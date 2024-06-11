package test

import (
	"fmt"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/role"
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
