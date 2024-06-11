package test

import (
	"fmt"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/commodity"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
)

func TestAddFavorite(t *testing.T) {
	conf := mysql.DbConfig{}
	db, err := mysql.Connect(&conf)
	if err != nil {
		t.Fatalf("Failed to load DB config from viper: %v", err)
	}
	err = db.AutoMigrate(&mysql.UserFavorite{}, &mysql.Commodity{})

	var userID uint = 1
	var favoriteID uint = 1001
	err = commodity.AddFavorite(db, userID, favoriteID)
	require.NoError(t, err, "add favorite failed")
	err = commodity.AddFavorite(db, userID, favoriteID)
	require.Error(t, err, "favorite should have be added")
	err = commodity.AddFavorite(db, userID, 999999)
	require.Error(t, err, "commodity doesn't exist")
}

func TestRemoveFavorite(t *testing.T) {
	conf := mysql.DbConfig{}
	db, err := mysql.Connect(&conf)
	if err != nil {
		t.Fatalf("Failed to load DB config from viper: %v", err)
	}
	err = db.AutoMigrate(&mysql.UserFavorite{})

	var userID uint = 1
	var favoriteID uint = 1001
	result := commodity.RemoveFavorite(db, userID, favoriteID)
	require.True(t, result, "remove favorite failed")
	result = commodity.RemoveFavorite(db, userID, favoriteID)
	require.False(t, result, "favorite should have be removed")
}

func TestGetFavorites(t *testing.T) {
	conf := mysql.DbConfig{}
	db, _ := mysql.Connect(&conf)
	_ = db.AutoMigrate(&mysql.Commodity{}, &mysql.UserFavorite{})

	testCommodities := []mysql.Commodity{
		{Model: gorm.Model{ID: 1}, Name: "test1", Price: 400},
		{Model: gorm.Model{ID: 2}, Name: "test2", Price: 200},
		{Model: gorm.Model{ID: 3}, Name: "test3", Price: 300},
		{Model: gorm.Model{ID: 4}, Name: "test4", Price: 100},
	}
	db.Create(&testCommodities)
	defer func() { db.Unscoped().Delete(&testCommodities) }()

	testUserFavorites := []mysql.UserFavorite{
		{UserID: 1, CommodityID: 1},
		{UserID: 1, CommodityID: 2},
		{UserID: 1, CommodityID: 3},
		{UserID: 1, CommodityID: 4},
		{UserID: 2, CommodityID: 2},
		{UserID: 2, CommodityID: 1},
	}
	db.Create(&testUserFavorites)
	defer func() { db.Unscoped().Delete(&testUserFavorites) }()

	//user1，price升序
	user1AscByPrice := commodity.GetFavorites(db, 1, "price", false, 1, 10)
	require.Equal(t, 4, len(user1AscByPrice), "user1AscByPrice should have four items")
	fmt.Println("user1's favorites, asc by price:")
	for _, _commodity := range user1AscByPrice {
		fmt.Printf("ID:%d\tPrice:%f\tName:%v\n", _commodity.ID, _commodity.Price, _commodity.Name)
	}

	//user2, created_at降序
	user2DescByCreatedAt := commodity.GetFavorites(db, 2, "created_at", true, 1, 10)
	require.Equal(t, 2, len(user2DescByCreatedAt), "user2DescByCreatedAt should have two items")
	fmt.Println("user2's favorites, desc by created_at:")
	for _, _commodity := range user2DescByCreatedAt {
		fmt.Printf("ID:%d\tCreate%v\tName:%v\n", _commodity.ID, _commodity.CreatedAt, _commodity.Name)
	}

	//user1, 默认降序分页
	user1DescByDefault := commodity.GetFavorites(db, 1, "", true, 2, 2)
	require.Equal(t, 2, len(user1DescByDefault), "user1DescByDefault should have two item")
	fmt.Println("user1's favorites, desc by default:")
	for _, _commodity := range user1DescByDefault {
		fmt.Printf("ID:%d\tName:%v\n", _commodity.ID, _commodity.Name)
	}
}
