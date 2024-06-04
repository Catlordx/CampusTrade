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
//	@return	bool		添加结果
func AddFavorite(db *gorm.DB, userID, favoriteID uint) bool {
	//判断商品是否已经被收藏
	result := db.Model(mysql.UserFavorite{}).
		Where("user_id = ? AND commodity_id = ?", userID, favoriteID).
		First(&mysql.UserFavorite{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) == false {
		return false
	}

	favorite := mysql.UserFavorite{
		UserID:      userID,
		CommodityID: favoriteID,
	}
	db.Create(&favorite)
	return true
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

// GetFavoriteIDs
//
//	@Description: 获取用户收藏商品ID的分页列表
//	@param	db		数据库DB
//	@param	kind	商品种类
//	@param	page	页数
//	@param	count	每页商品数
//	@return	[]uint	商品ID
func GetFavoriteIDs(db *gorm.DB, userID uint, page, count int) []uint {
	var commodityIDs []uint

	offset := (page - 1) * count
	//获取商品ID切片
	db.Model(&mysql.UserFavorite{}).
		Where("user_id = ?", userID).
		Limit(count).
		Offset(offset).
		Pluck("commodity_id", &commodityIDs)

	return commodityIDs
}
