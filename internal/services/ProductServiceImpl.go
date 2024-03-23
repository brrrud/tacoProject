package services

import (
	"tacoProject/internal/models"
	"tacoProject/internal/repositories"
)

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository}
}

func (p *ProductServiceImpl) CreateProduct(product models.ProductModel) error {
	err := p.ProductRepository.CreateProduct(product)
	return err
}
