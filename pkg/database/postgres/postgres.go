package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/t0nyandre/go-rest-boilerplate/internal/config"
	"go.uber.org/zap"
)

func NewPostgres(logger *zap.SugaredLogger, cfg *config.Config) (*sqlx.DB, error) {
	db, _ := sqlx.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresDb,
			cfg.PostgresSslMode))

	if err := db.DB.Ping(); err != nil {
		logger.Warnw(
			"Retrying database connection in 5 seconds",
			"appName", cfg.AppName,
			"user", cfg.PostgresUser,
			"database", cfg.PostgresDb,
			"error", err,
		)
		time.Sleep(time.Duration(5) * time.Second)
		return NewPostgres(logger, cfg)
	}

	logger.Infow(
		"Successfully connected to database",
		"appName", cfg.AppName,
		"user", cfg.PostgresUser,
		"database", cfg.PostgresDb,
	)

	return db, nil
}
