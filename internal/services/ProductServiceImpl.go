package services

import (
	"tacoProject/internal/models"
	"tacoProject/internal/repositories"
)

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepositoryImpl
}

func (p *ProductServiceImpl) CreateProduct(product models.ProductModel) error {
	err := p.ProductRepository.CreateProduct(product)
	return err
}
