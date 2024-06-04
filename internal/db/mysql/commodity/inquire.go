package commodity

import (
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"gorm.io/gorm"
)

// GetCommoditiesByID
//
//	@Description: 根据商品ID获取商品结构体切片
//	@param	db					数据库DB
//	@param	commodityIDs		商品ID切片
//	@return	[]mysql.Commodity	商品结构体切片
func GetCommoditiesByID(db *gorm.DB, commodityIDs []uint) []mysql.Commodity {
	var commodities []mysql.Commodity

	db.Model(mysql.Commodity{}).
		//Where("deleted_at IS NOT NULL").
		Where("id IN (?)", commodityIDs).
		Find(&commodities)

	return commodities
}

// GetCommodityIDsByKind
//
//	@Description: 获取kind种类商品ID的分页列表
//	@param	db		数据库DB
//	@param	kind	商品种类
//	@param	page	页数
//	@param	count	每页商品数
//	@return	[]uint	商品ID
func GetCommodityIDsByKind(db *gorm.DB, kind string, page, count int) []uint {
	var commodityIDs []uint

	offset := (page - 1) * count
	//kind为空时返回所有商品的分页列表
	if kind == "" {
		db.Model(&mysql.Commodity{}).
			//Where("deleted_at IS NOT NULL").
			Limit(count).
			Offset(offset).
			Pluck("id", &commodityIDs)
	} else {
		//获取商品ID切片
		db.Model(&mysql.CommodityKind{}).
			Where("kind = ?", kind).
			Limit(count).
			Offset(offset).
			Pluck("commodity_id", &commodityIDs)
	}

	return commodityIDs
}
