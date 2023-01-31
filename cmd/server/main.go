package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/t0nyandre/go-rest-boilerplate/internal/config"
	"github.com/t0nyandre/go-rest-boilerplate/internal/healthcheck"
	"github.com/t0nyandre/go-rest-boilerplate/internal/user"
	"github.com/t0nyandre/go-rest-boilerplate/pkg/database/postgres"
	"github.com/t0nyandre/go-rest-boilerplate/pkg/logger"
)

var appConfig = flag.String("config", "./config/local.yml", "path to config file")

func main() {
	flag.Parse()

	logger := logger.NewLogger()

	cfg, err := config.Load(*appConfig, logger)
	if err != nil {
		logger.Fatalw("Failed to load config", "error", err)
	}

	// Connect to database
	db, err := postgres.NewPostgres(logger, cfg)
	if err != nil {
		logger.Fatalw("Failed to connect to database", "database", cfg.PostgresDb, "error", err)
	}

	router := chi.NewRouter()
	router.Mount("/healthcheck", healthcheck.RegisterHandlers(cfg))
	router.Mount("/v1/users", user.RegisterHandlers(user.NewService(user.NewRepository(db, logger), logger), logger))

	logger.Infow("Server successfully up and running", "host", cfg.AppHost, "port", cfg.AppPort)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%v", cfg.AppHost, cfg.AppPort), router); err != nil {
		logger.Fatalw("Server failed to start", "error", err)
	}
}
