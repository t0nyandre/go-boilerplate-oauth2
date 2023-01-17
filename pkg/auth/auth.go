package auth

import "time"

type AuthProvider struct {
	ProviderName   string    `json:"provider_name,omitempty"`
	ProviderUserID string    `json:"provider_user_id,omitempty"`
	AccessToken    string    `json:"access_token,omitempty"`
	RefreshToken   string    `json:"refresh_token,omitempty"`
	Expiry         time.Time `json:"expiry,omitempty"`
	UserID         string    `json:"user_id,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
