package product

import (
	"github.com/Catlordx/CampusTrade/internal/core/config"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/commodity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// CommodityListByLike
//
//	@Description: 获取拥有某类型的商品分页列表
//	@param	c	gin.Context
func CommodityListByLike(c *gin.Context) {
	kind := c.PostForm("kind")
	sort := c.PostForm("count")
	reverseStr := c.PostForm("reverse")
	page, _ := strconv.Atoi(c.PostForm("page"))
	count, _ := strconv.Atoi(c.PostForm("count"))

	var reverse bool
	if strings.ToLower(reverseStr) == "asc" || strings.ToLower(reverseStr) == "" { //默认升序
		reverse = false
	} else if strings.ToLower(reverseStr) == "desc" {
		reverse = true
	}

	appContext := c.MustGet("appContext").(*config.AppContext)
	commodityList :=
		commodity.GetCommoditiesByKind(
			appContext.DB,
			kind,
			sort,
			reverse,
			page,
			count)

	c.JSON(http.StatusOK, commodityList)
}
