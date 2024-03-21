package services

import "tacoProject/internal/models"

type ProductService interface {
	CreateProduct(product models.ProductModel) error
}
