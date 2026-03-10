package handler

import (
	"context"
	"crud-gin-mongodb/dto"
	"crud-gin-mongodb/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.ProductService
}

func NewHandlerProduct(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var dto dto.CreateProdReq
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	product, err := h.service.CreateService(ctx, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "sucsess create product",
		"data":    product,
	})

}

func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()
	prod, err := h.service.FindAllProduct(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "sucsess get all product",
		"datas":   prod,
	})
}

func (h *ProductHandler) FindByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()
	id := c.Param("id")
	product, err := h.service.FindByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	c.JSON(200, gin.H{
		"datas": product,
	})
}

func (h *ProductHandler) UpdateByID(c *gin.Context) {
	var dto dto.CreateProdReq
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	id := c.Param("id")
	prodUpdate, err := h.service.UpdateByID(ctx, id, dto)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "sucsess updated",
		"data":    prodUpdate,
	})
}
