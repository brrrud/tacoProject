package services

import "tacoProject/internal/models"

type TacoService interface {
	FindById(id int64) (models.TacoModel, error)
	CreateTacoByProducts() (models.TacoModel, error)
}
