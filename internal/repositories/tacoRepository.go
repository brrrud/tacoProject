package repositories

import "tacoProject/internal/models"

type TacoRepository interface {
	FindById(id int64) (models.TacoModel, error)
	CreateTacoByProducts() (models.TacoModel, error)
}
