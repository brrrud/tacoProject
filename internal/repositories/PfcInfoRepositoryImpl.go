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

func (p *PfcInfoRepositoryImpl) CreatePfcInfo(info models.PfcInfo) (int64, error) {
	queryForCreatePfcInfo := `
	INSERT INTO pfc_info (proteins, fats, carbohydrates)
	VALUES
	($1, $2, $3) RETURNING pfc_info.id_pfc_info
	`
	var returningID int64
	err := p.Db.QueryRow(context.Background(), queryForCreatePfcInfo,
		info.Proteins,
		info.Fats,
		info.Carbohydrates).Scan(&returningID)
	if err != nil {
		return -1, fmt.Errorf("problem with insert product %w", err)
	}
	return returningID, nil
}
