package services

import "tacoProject/internal/models"

type TacoService interface {
	FindById(id int64) (models.TacoModel, error)
	FindByName(nameTaco string) (models.TacoModel, error)
	CreateTacoByProducts(request models.RequestForCreateTaco) error
}
