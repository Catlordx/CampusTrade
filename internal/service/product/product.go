package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProductDetailsHandler(c *gin.Context) {
	productId := c.Param("productId")

	// TODO 商品查询逻辑

	if productId == "1" {
		c.JSON(http.StatusOK, gin.H{
			"message":   "成功获取商品信息",
			"productId": productId,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "商品信息未找到",
		})
	}
}
