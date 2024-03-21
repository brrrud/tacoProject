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
	queryForTaco :=
		` 
		 SELECT t.id_taco,
       t.name_taco,
       SUM(pr.weight_product) AS total_weight_product,
       SUM(p.proteins)        AS total_proteins,
       SUM(p.fats)            AS total_fats,
       SUM(p.carbohydrates)   AS total_carbohydrates
FROM taco t
         JOIN taco_product tp ON t.id_taco = tp.taco_id
         JOIN product pr ON tp.product_id = pr.id_product
         JOIN pfc_info p ON pr.pfc_info_id_fk = p.id_pfc_info
WHERE t.name_taco = $1
GROUP BY t.id_taco,
         t.name_taco;

		`

	var taco models.TacoModel
	row := t.Db.QueryRow(context.Background(), queryForTaco, nameTaco)
	err := row.Scan(&taco.IdTaco,
		&taco.NameTaco,
		&taco.WeightTaco,
		&taco.PfcTaco.Proteins,
		&taco.PfcTaco.Fats,
		&taco.PfcTaco.Carbohydrates)
	if err != nil {
		return models.TacoModel{}, err
	}
	queryForGetProducts :=
		`
SELECT p.name_product
FROM product p
    JOIN taco_product tp ON p.id_product = tp.product_id
WHERE taco_id = $1
		`
	rows, err := t.Db.Query(context.Background(), queryForGetProducts, taco.IdTaco)
	for rows.Next() {
		var currentProduct string
		err := rows.Scan(&currentProduct)
		if err != nil {
			//TODO: Нужно использовать свой тип ошибок в случае если тако существует,
			//но он пустой(не содержит продуктов)
			return models.TacoModel{}, err
		}
		taco.ProductTacoNames = append(taco.ProductTacoNames, currentProduct)
	}
	return taco, err
}

func (t *TacoRepositoryImpl) CreateTacoByProducts() (models.TacoModel, error) {
	return models.TacoModel{}, nil
}
