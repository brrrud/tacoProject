package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"tacoProject/internal/models"
)

type TacoRepositoryImpl struct {
	Db *pgx.Conn
}

func NewTacoRepository(Db *pgx.Conn) TacoRepository {
	return &TacoRepositoryImpl{Db: Db}
}

func (u *TacoRepositoryImpl) FindById(id int64) (models.TacoModel, error) {
	q := `
		SELECT id_taco, name_taco FROM taco WHERE id_taco = $1
	`
	var taco models.TacoModel
	row := u.Db.QueryRow(context.Background(), q)
	err := row.Scan(&taco.IdTaco, &taco.NameTaco)
	return taco, err
}

func (u *TacoRepositoryImpl) CreateTacoByProducts() (models.TacoModel, error) {
	// Парсинг входных данных
	var request models.RequestForCreateTaco
	// Начало транзакции
	tx, err := u.Db.Begin(context.Background())
	if err != nil {
		log.Fatalf("Could not begin transaction: %v", err)
		return models.TacoModel{}, err
	}

	// Создание тако
	var tacoID int
	err = tx.QueryRow(context.Background(),
		"INSERT INTO taco (name_taco) VALUES ($1) RETURNING id_taco",
		request.NameTaco).Scan(&tacoID)

	if err != nil {
		err := tx.Rollback(context.Background())
		if err != nil {
			return models.TacoModel{}, err
		}
		log.Fatalf("Could not insert taco: %v", err)
	}

	// Добавление продуктов к тако
	for _, productID := range request.Products {
		// Проверка наличия продукта в базе
		var count int
		err = tx.QueryRow(context.Background(),
			"SELECT COUNT(*) FROM product WHERE id_product = $1",
			productID).Scan(&count)
		if err != nil {
			tx.Rollback(context.Background())
			log.Fatalf("Could not check product existence: %v", err)
			return models.TacoModel{}, err
		}
		if count == 0 {
			tx.Rollback(context.Background())
			//c.JSON(http.StatusBadRequest, gin.H{"error": "Product with ID " + string(productID) + " does not exist"})
			return models.TacoModel{}, err
		}

		_, err = tx.Exec(context.Background(),
			"INSERT INTO taco_product (taco_id, product_id) VALUES ($1, $2)",
			tacoID, productID)
		if err != nil {
			tx.Rollback(context.Background())
			log.Fatalf("Could not insert product for taco: %v", err)
		}
	}

	// Подтверждение транзакции
	err = tx.Commit(context.Background())
	if err != nil {
		log.Fatalf("Could not commit transaction: %v", err)
		return models.TacoModel{}, err
	}

	return models.TacoModel{}, nil
	//c.JSON(http.StatusCreated, gin.H{"message": "Taco created successfully"})
}
