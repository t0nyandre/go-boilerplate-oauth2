CONFIG_FILE ?= config/local.json
DB_USER := $(shell sed -n 's|.*"postgres_user": *"\([^"]*\)".*|\1|p' $(CONFIG_FILE))
DB_PASSWORD := $(shell sed -n 's|.*"postgres_password": *"\([^"]*\)".*|\1|p' $(CONFIG_FILE))
DB_HOST := $(shell sed -n 's|.*"postgres_host": *"\([^"]*\)".*|\1|p' $(CONFIG_FILE))
DB_PORT := $(shell sed -n 's|.*"postgres_port": *"\([^"]*\)".*|\1|p' $(CONFIG_FILE))
DB_SSL := $(shell sed -n 's|.*"postgres_ssl_mode": *"\([^"]*\)".*|\1|p' $(CONFIG_FILE))
DB_DATABASE := $(shell sed -n 's|.*"postgres_db": *"\([^"]*\)".*|\1|p' $(CONFIG_FILE))
MIGRATE := docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations -database "postgres://$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=$(DB_SSL)&user=$(DB_USER)&password=$(DB_PASSWORD)"

.PHONY: migrate-new
migrate-new:
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir /migrations $${name// /_}
	
.PHONY: migrate
migrate:
	@echo "Running all new database migrations ..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down:
	@echo "Reverting database to the last migration ..."
	@$(MIGRATE) down 1

.PHONY: migrate-reset
migrate-reset:
	@echo "Resetting database ..."
	@$(MIGRATE) drop -f
	@echo "Running all database migrations ..."
	@$(MIGRATE) up
