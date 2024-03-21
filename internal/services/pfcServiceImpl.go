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

//func NewPfcInfoService(pfcInfoRepository repositories.PfcInfoRepository) PfcInfoService {
//	return &PfcInfoService{
//		PfcInfoRepository: pfcIn,
//	}
//}

func (p *PfcInfoServiceImpl) CreateProduct(product models.PfcInfo) error {
	err := p.PfcInfoRepository.CreatePfcInfo(product)
	return err
}
