package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"tacoProject/internal/models"
)

type TacoRepositoryImpl struct {
	Db          *pgx.Conn
	productRepo ProductRepository
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
	ctx := context.Background()
	tx, errTr := t.Db.Begin(ctx)
	if errTr != nil {
		return models.TacoModel{}, fmt.Errorf("couldn't begin transaction: %w", errTr)
	}
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
			_ = tx.Rollback(ctx)
			return models.TacoModel{}, err
		}
		taco.ProductTacoNames = append(taco.ProductTacoNames, currentProduct)
	}
	return taco, err
}

func (t *TacoRepositoryImpl) CreateTacoByProducts(request models.RequestForCreateTaco) error {
	ctx := context.Background()
	tx, err := t.Db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("couldn't begin transaction: %w", err)
	}

	queryForCreateTaco := `
	INSERT INTO taco (name_taco) 
	VALUES ($1)
	RETURNING taco.id_taco
		`

	var tacoId int
	err = t.Db.QueryRow(context.Background(), queryForCreateTaco, request.NameTaco).Scan(&tacoId)
	if err != nil {
		return fmt.Errorf("create taco with product failed :\n%w", err)
	}

	queryForAddProducts := `
	INSERT INTO taco_product
	VALUES ($1, $2)
	`

	queryForCheckProductAvail := `
	SELECT product.id_product
	FROM product 
	WHERE id_product = ($1)
	`

	var exist int
	for _, v := range request.Products {
		err1 := t.Db.QueryRow(context.Background(), queryForCheckProductAvail, v).Scan(&exist)
		if err1 != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("no such product for taco: \n%w", err1)
		}

		_, err = t.Db.Exec(context.Background(), queryForAddProducts, tacoId, v)
		if err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("failed to add product into taco: \n%w", err)
		}
	}
	if err := tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		return fmt.Errorf("couldn't commit transaction: \n%w", err)
	}
	return nil

}
