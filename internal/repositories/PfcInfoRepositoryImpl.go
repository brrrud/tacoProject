package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"tacoProject/internal/models"
)

type PfcInfoRepositoryImpl struct {
	Db *pgx.Conn
}

func NewPfcInfoRepository(Db *pgx.Conn) PfcInfoRepository {
	return &PfcInfoRepositoryImpl{Db: Db}
}

func (p *PfcInfoRepositoryImpl) CreatePfcInfo(info models.PfcInfo) error {
	queryForCreatePfcInfo := `
	INSERT INTO pfc_info (proteins, fats, carbohydrates)
	VALUES
	($1, $2, $3)
	`
	_, err := p.Db.Exec(context.Background(), queryForCreatePfcInfo,
		info.Proteins,
		info.Fats,
		info.Carbohydrates)
	if err != nil {
		return fmt.Errorf("problem with insert product %w", err)
	}
	return nil
}
