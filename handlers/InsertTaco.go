package handlers

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type TacoRepositoryImpl struct {
	Db *pgx.Conn
}

func (pg *TacoRepositoryImpl) InsertUser(ctx context.Context) error {
	query := `INSERT INTO taco (name_taco) VALUES (@taco)`
	args := pgx.NamedArgs{
		"userName":  "Bobby",
		"userEmail": "bobby@donchev.is",
	}

	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
