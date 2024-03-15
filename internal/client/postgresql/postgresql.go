package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func OpenDbConnection() (*pgx.Conn, error) {
	ctx := context.Background()

	db, err := pgx.Connect(ctx, os.Getenv("DB_SERVER_URL"))
	if err != nil {
		return nil, fmt.Errorf("Failed, not connected to database, %v", err)
	}

	return db, nil
}

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
