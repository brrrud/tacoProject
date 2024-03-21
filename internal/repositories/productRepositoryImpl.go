package repositories

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"tacoProject/internal/models"
)

type ProductRepositoryImpl struct {
	Db *pgx.Conn
}

func (p *ProductRepositoryImpl) CreateProduct(product models.ProductModel) error {
	queryForCreateProduct :=
		`
			INSERT INTO product (name_product, pfc_info_id_fk, weight_product)
			VALUES 
			($1, $2, $3)                                                              
		`
	fmt.Println(queryForCreateProduct)
	return nil
}
