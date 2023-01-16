package auth

import "github.com/jmoiron/sqlx"

type AuthRepository interface {
	Create(auth *AuthProvider) (*AuthProvider, error)
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(auth *AuthProvider) (*AuthProvider, error) {
	query := `
        INSERT INTO auth (provider_name, provider_user_id, access_token, refresh_token, expiry, user_id)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, provider_name, provider_user_id, access_token, refresh_token, expiry, user_id
    `
	err := r.db.QueryRowx(query, auth.ProviderName, auth.ProviderUserID, auth.AccessToken, auth.RefreshToken, auth.Expiry, auth.UserID).StructScan(auth)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
