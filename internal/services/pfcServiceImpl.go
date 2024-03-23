package services

import (
	"tacoProject/internal/models"
	"tacoProject/internal/repositories"
)

type PfcInfoServiceImpl struct {
	PfcInfoRepository repositories.PfcInfoRepository
}

func NewPfcInfoServiceImpl(pfcInfoRepository repositories.PfcInfoRepository) PfcInfoService {
	return &PfcInfoServiceImpl{PfcInfoRepository: pfcInfoRepository}
}

func (p *PfcInfoServiceImpl) CreateProductPFC(product models.PfcInfo) (int64, error) {
	pfcInfoId, err := p.PfcInfoRepository.CreatePfcInfo(product)
	return pfcInfoId, err
}
