package test

import (
	"fmt"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/commodity"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommodityIDList(t *testing.T) {
	conf := mysql.DbConfig{}
	db, err := mysql.Connect(&conf)
	if err != nil {
		t.Fatalf("Failed to load DB config from viper: %v", err)
	}

	err = db.AutoMigrate(&mysql.Commodity{}, &mysql.CommodityKind{})
	testComm1 := &mysql.Commodity{
		Name:        "test1",
		Description: "the first commodity to be tested",
	}
	testComm2 := &mysql.Commodity{
		Name:        "test2",
		Description: "the second commodity to be tested",
	}
	db.Create(&testComm1)
	db.Create(&testComm2)
	defer db.Delete(&testComm1)
	defer db.Delete(&testComm2)

	comm1Kind1 := mysql.CommodityKind{
		CommodityID: testComm1.ID,
		Kind:        "testKind1",
	}
	db.Create(&comm1Kind1)
	defer db.Delete(&comm1Kind1)

	commodities1 := commodity.GetCommodityIDsByKind(db, "", 1, 10)
	require.NotZero(t, len(commodities1), "expected commodities1 to be non-empty")

	commodities2 := commodity.GetCommodityIDsByKind(db, "", 2, 1)
	require.Len(t, commodities2, 1, "commodities2 should have 1 element")

	commoditiesKind1 := commodity.GetCommodityIDsByKind(db, "testKind1", 1, 10)
	require.Len(t, commoditiesKind1, 1, "commoditiesKind1 should have 1 element")

	commoditiesKind2 := commodity.GetCommodityIDsByKind(db, "testKind2", 1, 10)
	require.Len(t, commoditiesKind2, 0, "commoditiesKind2 should have 0 elements")

}

func TestCommodityListByID(t *testing.T) {
	conf := mysql.DbConfig{}
	db, err := mysql.Connect(&conf)
	if err != nil {
		t.Fatalf("Failed to load DB config from viper: %v", err)
	}
	err = db.AutoMigrate(&mysql.Commodity{}, &mysql.CommodityKind{})

	commodityIDs1 := commodity.GetCommodityIDsByKind(db, "", 1, 10)
	require.NotZero(t, len(commodityIDs1), "expected commodities1 to be non-empty")

	commodities1 := commodity.GetCommoditiesByID(db, commodityIDs1)
	require.NotZero(t, len(commodities1), "expected commodities1 to be non-empty")

	for id, _commodity := range commodities1 {
		fmt.Printf("%d %v\n", id, _commodity.Name)
	}
}
