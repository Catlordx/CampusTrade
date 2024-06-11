package commodity

import (
	"errors"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"gorm.io/gorm"
	"strings"
)

type CommoditiesInfo struct {
	Kind    string `json:"kind"`    //列表按照kind进行分类，kind为空字符串时表示不分类，即从所有商品中寻找
	Sort    string `json:"sort"`    //列表按照sort的值进行排序，sort为空字符串时默认按照ID排序
	Reverse string `json:"reverse"` //排序规则，ASC升序，DESC降序
	Page    int    `json:"page"`    //第page页
	Count   int    `json:"count"`   //每页商品个数
}

// GetCommodityByID
//
//	@Description:  根据ID查询商品
//	@param	db					数据库DB
//	@param	ID					查询商品ID
//	@return	*mysql.Commodity	商品结构体
func GetCommodityByID(db *gorm.DB, ID uint) *mysql.Commodity {
	var commodity mysql.Commodity
	result := db.First(&commodity, ID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &commodity
}

// GetCommoditiesByKind
//
//	@Description: 获取商品分类分页排序列表
//	@param	db					数据库DB
//	@param	listInfo			商品列表的限制信息
//	@return	[]mysql.Commodity	商品结构体列表
func GetCommoditiesByKind(db *gorm.DB, listInfo CommoditiesInfo) []mysql.Commodity {
	var commodities []mysql.Commodity

	var order string
	// 默认按照ID排序
	if listInfo.Sort == "" {
		order = "id"
	} else {
		order = listInfo.Sort
	}
	// 根据 reverse 获取排序规则
	if listInfo.Reverse == "" || strings.ToUpper(listInfo.Reverse) == "ASC" { //升序
		order += " ASC"
	} else { // 降序
		order += " DESC"
	}

	// 构建查询
	query := db.Model(&mysql.Commodity{})

	// 如果 kind 为空，从所有商品中寻找
	// 如果 kind 不为空，则加入筛选条件
	if listInfo.Kind != "" {
		subQuery :=
			db.Model(&mysql.CommodityKind{}).
				Select("commodity_id").
				Where("kind = ?", listInfo.Kind)
		query = query.Where("id IN (?)", subQuery)
	}

	offset := (listInfo.Page - 1) * listInfo.Count
	// 查询 Commodity 表，获取指定页数和数量的记录，并按排序参数排序
	query.Order(order).
		Offset(offset).
		Limit(listInfo.Count).
		Find(&commodities)
	return commodities
}
