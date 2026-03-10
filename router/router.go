package router

import (
	"crud-gin-mongodb/handler"

	"github.com/gin-gonic/gin"
)

func Router(handlerProduct handler.ProductHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/v1/api")
	{
		product := api.Group("/product")
		{
			product.POST("/create", handlerProduct.CreateProduct)
			product.GET("/", handlerProduct.GetAllProduct)
			product.GET("/:id", handlerProduct.FindByID)
			product.POST("/:id", handlerProduct.UpdateByID)
		}
	}
	return r
}
