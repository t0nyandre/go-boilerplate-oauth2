package user

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lithammer/shortuuid"
)

type UserRepository interface {
	Create(user *User) (*User, error)
}

type repository struct {
	db *sqlx.DB
}

// Create implements UserRepository
func (r *repository) Create(user *User) (*User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.ID = shortuuid.New()

	query := `
        INSERT INTO users (id, name, username, email, is_active, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, username, email, is_active, created_at
    `
	err := r.db.QueryRowx(query, user.ID, user.Name, user.Username, user.Email, user.IsActive, user.CreatedAt, user.UpdatedAt).StructScan(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewRepository(db *sqlx.DB) UserRepository {
	return &repository{db: db}
}
