package test

import (
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
			t.Errorf("User already exists")
		} else {
			t.Errorf("password should be valid")
		}
	}

	err = user.AddUser(db,
		testUser1Username,
		testUser1RealName,
		testUser1Password,
		testUser1PhoneNumber,
		testUser1Role)
	require.Error(t, err, "An error was expected but got nil")

	err = user.AddUser(db,
		"test_user2",
		"",
		"123456",
		"",
		role.User)
	require.NoError(t, err, "an error should be nil")

	err = user.AddUser(db, "", "", "", "", "")
	require.Error(t, err, "An error was expected but got nil")
	err = user.AddUser(db, "test_user3", "", "", "", "")
	require.Error(t, err, "An error was expected but got nil")
	err = user.AddUser(db, "test_user4", "", "123456", "", "")
	require.Error(t, err, "An error was expected but got nil")
}

func TestModifyUser(t *testing.T) {
	conf := mysql.DbConfig{}
	db, _ := mysql.Connect(&conf)

	_ = db.AutoMigrate(&mysql.User{})

	testUser1 := user.GetUserByUsername(db, "test_user1")
	testUser1id := testUser1.ID
	testUser1Username := testUser1.Username
	testUser1RealName := testUser1.RealName
	testUser1Password := testUser1.Password
	testUser1PhoneNumber := testUser1.PhoneNumber

	user.ModifyUser(db, testUser1Username, "", "", "", "")
	_testUser1 := user.GetUserByID(db, testUser1id)
	require.Equal(t, _testUser1.Username, testUser1Username, "test_user1's username shouldn't be changed")
	require.Equal(t, _testUser1.RealName, testUser1RealName, "test_user1's real name shouldn't be changed")
	require.Equal(t, string(_testUser1.Password), string(testUser1Password), "test_user1's password shouldn't changed")
	require.Equal(t, _testUser1.PhoneNumber, testUser1PhoneNumber, "test_user1's phone number shouldn't be changed")

	user.ModifyUser(db,
		"test_user1",
		"test_user1.0",
		"TestUser1",
		"abcdef",
		"12345678910")
	testUser10 := user.GetUserByID(db, testUser1id)

	require.Equal(t, testUser10.Username, "test_user1.0", "test_user1's username doesn't be changed")
	require.Equal(t, testUser10.RealName, "TestUser1", "test_user1's real name doesn't be changed")
	require.True(t, user.CheckPassword(testUser10.Password, "abcdef"), "test_user1's password doesn't be changed")
	require.Equal(t, testUser10.PhoneNumber, "12345678910", "test_user1's phone number doesn't be changed")
	user.ModifyUser(db,
		testUser10.Username,
		testUser1Username,
		testUser1RealName,
		"123456",
		testUser1PhoneNumber)

	result := user.ModifyUser(db, "", "", "", "", "")
	require.False(t, result, "username is empty, should not be changed")
	result = user.ModifyUser(db, "no_user", "", "", "", "")
	require.False(t, result, "username is empty, should not be changed")
}
