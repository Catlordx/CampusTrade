package router

import (
	_product "github.com/Catlordx/CampusTrade/internal/service/product"
	"github.com/gin-gonic/gin"
)

func setProductRouter(r *gin.Engine) {
	productGroup := r.Group("/products")
	{
		productGroup.POST("/")
		productGroup.GET("/:productId", _product.GetProductDetailsHandler)
		productGroup.GET("/user/:userId")
		productGroup.PUT("/:productId")

		// TODO 获取商品分页列表
		productGroup.GET("/list/:kind/:sort/:reverse/:page/:count", _product.CommodityListByLike)
	}
}
