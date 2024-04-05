package repositories

import "tacoProject/internal/models"

type TacoRepository interface {
	FindById(id int64) (models.TacoModel, error)
	FindByName(nameTaco string) (models.TacoModel, error)
	CreateTacoByProducts(request models.RequestForCreateTaco) error
	GetTacosLessThanPFC(pfcSum float64) ([]models.TacoModel, error)
}
