package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"tacoProject/internal/models"
)

type TacoRepositoryImpl struct {
	Db *pgx.Conn
}

func NewTacoRepository(Db *pgx.Conn) TacoRepository {
	return &TacoRepositoryImpl{Db: Db}
}

func (t *TacoRepositoryImpl) FindById(id int64) (models.TacoModel, error) {
	q := `SELECT id_taco, name_taco  FROM taco WHERE id_taco = $1`

	var taco models.TacoModel
	row := t.Db.QueryRow(context.Background(), q, id)
	err := row.Scan(&taco.IdTaco, &taco.NameTaco)
	return taco, err
}

func (t *TacoRepositoryImpl) FindByName(nameTaco string) (models.TacoModel, error) {
	q := "SELECT id_taco, name_taco FROM taco where name_taco = $1"
	var taco models.TacoModel
	row := t.Db.QueryRow(context.Background(), q, nameTaco)
	err := row.Scan(&taco.IdTaco, &taco.NameTaco)
	return taco, err
}

func (t *TacoRepositoryImpl) CreateTacoByProducts() (models.TacoModel, error) {
	return models.TacoModel{}, nil
}
