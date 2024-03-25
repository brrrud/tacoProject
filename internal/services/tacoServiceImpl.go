package services

import (
	"tacoProject/internal/models"
	"tacoProject/internal/repositories"
)

type TacoServiceImpl struct {
	TacoRepository repositories.TacoRepository
}

func (t *TacoServiceImpl) CreateTacoByProducts(request models.RequestForCreateTaco) error {
	err := t.TacoRepository.CreateTacoByProducts(request)
	return err
}

func (t *TacoServiceImpl) FindById(id int64) (models.TacoModel, error) {
	taco, err := t.TacoRepository.FindById(id)
	return taco, err
}
func (t *TacoServiceImpl) FindByName(nameTaco string) (models.TacoModel, error) {
	taco, err := t.TacoRepository.FindByName(nameTaco)
	return taco, err
}

func NewTacoServiceImpl(tacoRepository repositories.TacoRepository) TacoService {
	return &TacoServiceImpl{
		TacoRepository: tacoRepository,
	}
}
