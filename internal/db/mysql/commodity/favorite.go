package commodity

import (
	"errors"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"gorm.io/gorm"
)

// AddFavorite
//
//	@Description: 添加收藏商品
//	@param	db			数据库DB
//	@param	userID		用户UI
//	@param	favoriteID	收藏商品UI
//	@return	error		添加结果
func AddFavorite(db *gorm.DB, userID, favoriteID uint) error {
	//检查商品是否存在
	commodity := GetCommodityByID(db, favoriteID)
	if commodity == nil {
		return errors.New("commodity not found")
	}
	//判断商品是否已经被收藏
	result := db.Model(mysql.UserFavorite{}).
		Where("user_id = ? AND commodity_id = ?", userID, favoriteID).
		First(&mysql.UserFavorite{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) == false {
		return errors.New("commodity already be in favorite")
	}

	favorite := mysql.UserFavorite{
		UserID:      userID,
		CommodityID: favoriteID,
	}
	db.Create(&favorite)
	return nil
}

// RemoveFavorite
//
//	 @Description: 移除收藏商品
//		@param	db			数据库DB
//		@param	userID		用户UI
//		@param	favoriteID	收藏商品UI
//		@return	bool		移除结果
func RemoveFavorite(db *gorm.DB, userID, favoriteID uint) bool {
	var userFavorite mysql.UserFavorite
	result := db.Model(mysql.UserFavorite{}).
		Where("user_id = ? AND commodity_id = ?", userID, favoriteID).
		First(&userFavorite)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	db.Delete(&userFavorite)
	return true
}

// GetFavorites
//
//	@Description: 获取收藏商品分类分页列表，该列表按照sort的值进行排序
//	@param	db					数据库DB
//	@param	userID				用户ID
//	@param	sort				列表按照sort的值进行排序，sort为空字符串时默认按照ID排序
//	@param	reverse				排序规则，true降序排列，false升序排列
//	@param	page				第page页
//	@param	count				每页商品个数
//	@return	[]mysql.Commodity	商品列表
func GetFavorites(db *gorm.DB, userID uint, sort string, reverse bool, page, count int) []mysql.Commodity {
	var commodities []mysql.Commodity

	var order string
	// 默认按照ID排序
	if sort == "" {
		order = "id"
	} else {
		order = sort
	}
	// 根据 reverse 获取排序规则
	if reverse { // 降序
		order += " DESC"
	} else { //升序
		order += " ASC"
	}

	subQuery :=
		db.Model(&mysql.UserFavorite{}).
			Select("commodity_id").
			Where("user_id = ?", userID)

	offset := (page - 1) * count
	db.Model(&mysql.Commodity{}).
		Where("id IN (?)", subQuery).
		Order(order).
		Offset(offset).
		Limit(count).
		Find(&commodities)

	return commodities
}
