package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tacoProject/internal/models"
	"tacoProject/internal/services"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (controller *ProductController) PostProduct(ctx *gin.Context) {
	var newProduct models.ProductModel
	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := controller.productService.CreateProduct(newProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create product"})
		return
	}
	ctx.Status(http.StatusOK)
}
