package commodity

import (
	"errors"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"gorm.io/gorm"
	"strings"
)

type FavoritesInfo struct {
	UserID  uint   //用户ID，列表查询ID为UserID的用户收藏的商品
	Sort    string `json:"sort"`    //列表按照sort的值进行排序，sort为空字符串时默认按照ID排序
	Reverse string `json:"reverse"` //排序规则，ASC升序，DESC降序
	Page    int    `json:"page"`    //第page页
	Count   int    `json:"count"`   //每页商品个数
}

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
//	@Description: 获取收藏商品列表
//	@param	db					数据库DB
//	@param	listInfo			收藏商品列表的限制信息
//	@return	[]mysql.Commodity	收藏商品列表
func GetFavorites(db *gorm.DB, listInfo FavoritesInfo) []mysql.Commodity {
	var favorites []mysql.Commodity

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

	subQuery :=
		db.Model(&mysql.UserFavorite{}).
			Select("commodity_id").
			Where("user_id = ?", listInfo.UserID)

	offset := (listInfo.Page - 1) * listInfo.Count
	db.Model(&mysql.Commodity{}).
		Where("id IN (?)", subQuery).
		Order(order).
		Offset(offset).
		Limit(listInfo.Count).
		Find(&favorites)

	return favorites
}
