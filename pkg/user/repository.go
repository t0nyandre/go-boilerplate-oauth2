package user

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(user *User) (*User, error)
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user *User) (*User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	query := `
        INSERT INTO users (name, username, email, is_active, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, username, email, is_active, created_at
    `
	err := r.db.QueryRowx(query, user.Name, user.Username, user.Email, user.IsActive, user.CreatedAt, user.UpdatedAt).StructScan(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
