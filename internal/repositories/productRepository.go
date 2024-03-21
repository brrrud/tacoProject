package repositories

import "tacoProject/internal/models"

type ProductRepository interface {
	CreateProduct(product models.ProductModel) error
}
