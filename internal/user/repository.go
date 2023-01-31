package user

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/t0nyandre/go-rest-boilerplate/internal/entity"
	"go.uber.org/zap"
)

type Repository interface {
	Create(user *entity.User) (string, error)
}

type repository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

// Create implements Repository
func (r repository) Create(user *entity.User) (string, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	query := `
        INSERT INTO users (id, name, username, password, email, is_active, created_at, updated_at)
        VALUES (:id, :name, :username, :password, :email, :is_active, :created_at, :updated_at)
    `
	_, err := r.db.NamedExec(query, &user)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func NewRepository(db *sqlx.DB, logger *zap.SugaredLogger) Repository {
	return repository{db: db, logger: logger}
}
