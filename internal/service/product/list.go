package product

import (
	"github.com/Catlordx/CampusTrade/internal/core/config"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/commodity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CommodityListByLike
//
//	@Description: 获取拥有某类型的商品分页列表
//	@param	c	gin.Context
func CommodityListByLike(c *gin.Context) {
	var commodityListInfo commodity.CommoditiesInfo
	err := c.ShouldBindJSON(&commodityListInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appContext := c.MustGet("appContext").(*config.AppContext)
	commodityList := commodity.GetCommoditiesByKind(appContext.DB, commodityListInfo)

	c.JSON(http.StatusOK, commodityList)
}
