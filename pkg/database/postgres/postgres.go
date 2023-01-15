package postgres

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func NewPostgres(logger *zap.SugaredLogger) (*sqlx.DB, error) {
	db, _ := sqlx.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_SSLMODE")))

	if err := db.DB.Ping(); err != nil {
		logger.Warnw(
			"Retrying database connection in 5 seconds",
			"appName", os.Getenv("APP_NAME"),
			"user", os.Getenv("POSTGRES_USER"),
			"database", os.Getenv("POSTGRES_DB"),
			"error", err,
		)
		time.Sleep(time.Duration(5) * time.Second)
		return NewPostgres(logger)
	}

	logger.Infow(
		"Successfully connected to database",
		"database", os.Getenv("POSTGRES_DB"),
	)

	return db, nil
}
