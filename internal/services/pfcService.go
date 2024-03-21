package services

import "tacoProject/internal/models"

type PfcInfoService interface {
	CreateProduct(product models.PfcInfo) error
}
