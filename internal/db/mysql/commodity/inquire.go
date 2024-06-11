package commodity

import (
	"errors"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"gorm.io/gorm"
)

/*
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
}*/

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
//	@Description: 获取商品分类分页列表，该列表按照sort的值进行排序
//	@param	db					数据库DB
//	@param	kind				列表按照kind进行分类，kind为空字符串时表示不分类，即从所有商品中寻找
//	@param	sort				列表按照sort的值进行排序，sort为空字符串时默认按照ID排序
//	@param	reverse				排序规则，true降序排列，false升序排列
//	@param	page				第page页
//	@param	count				每页商品个数
//	@return	[]mysql.Commodity	商品列表
func GetCommoditiesByKind(db *gorm.DB, kind, sort string, reverse bool, page, count int) []mysql.Commodity {
	var commodities []mysql.Commodity

	var order string
	// 默认按照ID排序
	if sort == "" {
		order = "id"
	} else {
		order = sort
	}
	// 根据 reverse 获取排序规则
	if reverse == false { //升序
		order += " ASC"
	} else { // 降序
		order += " DESC"
	}

	// 构建查询
	query := db.Model(&mysql.Commodity{})

	// 如果 kind 为空，从所有商品中寻找
	// 如果 kind 不为空，则加入筛选条件
	if kind != "" {
		subQuery :=
			db.Model(&mysql.CommodityKind{}).
				Select("commodity_id").
				Where("kind = ?", kind)
		query = query.Where("id IN (?)", subQuery)
	}

	offset := (page - 1) * count
	// 查询 Commodity 表，获取指定页数和数量的记录，并按排序参数排序
	query.Order(order).
		Offset(offset).
		Limit(count).
		Find(&commodities)
	return commodities
}
