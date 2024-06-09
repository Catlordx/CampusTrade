package test

import (
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/commodity"
	"github.com/stretchr/testify/require"
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
	db, err := mysql.Connect(&conf)
	if err != nil {
		t.Fatalf("Failed to load DB config from viper: %v", err)
	}
	err = db.AutoMigrate(&mysql.UserFavorite{})

	_ = commodity.AddFavorite(db, 1, 1001)
	_ = commodity.AddFavorite(db, 1, 1002)
	defer commodity.RemoveFavorite(db, 1, 1001)
	defer commodity.RemoveFavorite(db, 1, 1002)

	user1Favorites1 := commodity.GetFavoriteIDs(db, 1, 1, 10)
	require.NotZero(t, len(user1Favorites1), "expected commodities1 to be non-empty")
	user1Favorites2 := commodity.GetFavoriteIDs(db, 1, 2, 1)
	require.Len(t, user1Favorites2, 1, "commodities2 should have 1 element")
	user1Favorites3 := commodity.GetFavoriteIDs(db, 1, 9999, 10)
	require.Len(t, user1Favorites3, 0, "commodities2 should have 0 element")
}
