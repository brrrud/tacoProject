package controllers

import (
	"tacoProject/internal/services"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

//func (controller *ProductController) PostProduct(ctx gin.Context) {
//	//err := controller.productService.CreateProduct()
//}
