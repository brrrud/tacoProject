package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"tacoProject/internal/models"
)

type ProductRepositoryImpl struct {
	Db      *pgx.Conn
	pfcRepo PfcInfoRepository
}

func NewProductRepository(db *pgx.Conn, pfcRepo PfcInfoRepository) ProductRepository {
	return &ProductRepositoryImpl{Db: db, pfcRepo: pfcRepo}
}

func (p *ProductRepositoryImpl) CreateProduct(product models.ProductModel) error {
	ctx := context.Background()
	tx, err := p.Db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("couldn't begin transaction: %w", err)
	}

	var pfcId int64
	pfcId, err = p.pfcRepo.CreatePfcInfo(product.PfcProduct)

	if err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			return err
		}
		return fmt.Errorf("failed to create PfcInfo: %w", err)
	}
	fmt.Println(pfcId)
	queryForCreateProduct :=
		`
			INSERT INTO product (name_product, pfc_info_id_fk, weight_product)
			VALUES 
			($1, $2, $3)                                                              
		`
	_, err = p.Db.Exec(ctx, queryForCreateProduct,
		product.NameProduct,
		pfcId,
		product.WeightProduct,
	)
	if err != nil {
		_ = tx.Rollback(ctx)
		return fmt.Errorf("failed to create Product: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		return fmt.Errorf("couldn't commit transaction: %w", err)
	}

	return nil
}
