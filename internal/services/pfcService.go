package services

import "tacoProject/internal/models"

type PfcInfoService interface {
	CreateProductPFC(product models.PfcInfo) (int64, error)
}
