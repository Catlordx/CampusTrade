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
	Kind1AscByPrice :=
		commodity.GetCommoditiesByKind(
			db,
			commodity.CommoditiesInfo{
				Kind:    "kind1",
				Sort:    "price",
				Reverse: "ASC",
				Page:    1,
				Count:   10})
	require.Equal(t, 2, len(Kind1AscByPrice), "Kind1AscByPrice should have two items")
	fmt.Println("classified by kind1, asc by price:")
	for _, _commodity := range Kind1AscByPrice {
		fmt.Printf("ID:%d\tPrice:%f\tName:%v\n", _commodity.ID, _commodity.Price, _commodity.Name)
	}

	//kind2，默认降序
	Kind2DescByCreatedAt :=
		commodity.GetCommoditiesByKind(
			db,
			commodity.CommoditiesInfo{
				Kind:    "kind2",
				Sort:    "",
				Reverse: "DESC",
				Page:    1,
				Count:   10})
	require.Equal(t, 2, len(Kind2DescByCreatedAt), "Kind2DescByCreatedAt should have two items")
	fmt.Println("classified by kind2, desc by default:")
	for _, _commodity := range Kind2DescByCreatedAt {
		fmt.Printf("ID:%d\tName:%v\n", _commodity.ID, _commodity.Name)
	}

	//不分类，created_at降序分页
	NoKindDescByDefault :=
		commodity.GetCommoditiesByKind(
			db, commodity.CommoditiesInfo{
				Kind:    "",
				Sort:    "created_at",
				Reverse: "DESC",
				Page:    2,
				Count:   10})
	require.NotZero(t, len(NoKindDescByDefault), "NoKindDescByDefault should not be empty")
	fmt.Println("not classified, desc by created_at:")
	for _, _commodity := range NoKindDescByDefault {
		fmt.Printf("ID:%d\tName:%v\n", _commodity.ID, _commodity.Name)
	}
}
