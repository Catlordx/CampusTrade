package test

import (
	"fmt"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/commodity"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
)

func TestGetCommoditiesByKind(t *testing.T) {
	conf := mysql.DbConfig{}
	db, _ := mysql.Connect(&conf)
	_ = db.AutoMigrate(&mysql.Commodity{}, &mysql.CommodityKind{})

	testCommodities := []mysql.Commodity{
		{Model: gorm.Model{ID: 1}, Name: "test1", Price: 300},
		{Model: gorm.Model{ID: 2}, Name: "test2", Price: 100},
		{Model: gorm.Model{ID: 3}, Name: "test3", Price: 200},
	}
	db.Create(&testCommodities)
	defer func() { db.Unscoped().Delete(&testCommodities) }()

	testCommodityKinds := []mysql.CommodityKind{
		{CommodityID: 1, Kind: "kind1"},
		{CommodityID: 2, Kind: "kind1"},
		{CommodityID: 2, Kind: "kind2"},
		{CommodityID: 3, Kind: "kind2"},
	}
	db.Create(&testCommodityKinds)
	defer func() { db.Unscoped().Delete(&testCommodityKinds) }()

	//kind1，price升序
	Kind1AscByPrice := commodity.GetCommoditiesByKind(db, "kind1", "price", false, 1, 10)
	require.Equal(t, 2, len(Kind1AscByPrice), "Kind1AscByPrice should have two items")
	fmt.Println("classified by kind1, asc by price:")
	for _, _commodity := range Kind1AscByPrice {
		fmt.Printf("ID:%d\tPrice:%f\tName:%v\n", _commodity.ID, _commodity.Price, _commodity.Name)
	}

	//kind2，默认降序
	Kind2DescByCreatedAt := commodity.GetCommoditiesByKind(db, "kind2", "", true, 1, 10)
	require.Equal(t, 2, len(Kind2DescByCreatedAt), "Kind2DescByCreatedAt should have two items")
	fmt.Println("classified by kind2, desc by default:")
	for _, _commodity := range Kind2DescByCreatedAt {
		fmt.Printf("ID:%d\tName:%v\n", _commodity.ID, _commodity.Name)
	}

	//不分类，created_at降序分页
	NoKindDescByDefault := commodity.GetCommoditiesByKind(db, "", "created_at", true, 2, 10)
	require.NotZero(t, len(NoKindDescByDefault), "NoKindDescByDefault should not be empty")
	fmt.Println("not classified, desc by created_at:")
	for _, _commodity := range NoKindDescByDefault {
		fmt.Printf("ID:%d\tName:%v\n", _commodity.ID, _commodity.Name)
	}
}

/*
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

	var commodityIDs2 = []uint{10000, 10001}
	commodities2 := commodity.GetCommoditiesByID(db, commodityIDs2)
	require.Zero(t, len(commodities2), "expected commodities1 to be empty")

	for id, _commodity := range commodities2 {
		fmt.Printf("%d %v\n", id, _commodity.Name)
	}
}
*/
