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
	}
}
