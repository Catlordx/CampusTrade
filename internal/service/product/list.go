package product

import (
	"github.com/Catlordx/CampusTrade/internal/core/config"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/commodity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CommodityListByLike
//
//	@Description: 获取拥有某类型的商品分页列表
//	@param	c	gin.Context
func CommodityListByLike(c *gin.Context) {
	like := c.PostForm("like")
	page, _ := strconv.Atoi(c.PostForm("page"))
	count, _ := strconv.Atoi(c.PostForm("count"))

	appContext := c.MustGet("appContext").(*config.AppContext)
	commodityList :=
		commodity.GetCommoditiesByID(
			appContext.DB,
			commodity.GetCommodityIDsByKind(
				appContext.DB,
				like,
				page,
				count,
			),
		)

	c.JSON(http.StatusOK, commodityList)
}
