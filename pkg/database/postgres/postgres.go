package postgres

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func NewPostgres(logger *zap.SugaredLogger) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_SSLMODE")))
	if err != nil {
		return nil, err
	}

	logger.Infow("Successfully connected to database",
		"database", os.Getenv("POSTGRES_DB"))

	return db, nil
}
