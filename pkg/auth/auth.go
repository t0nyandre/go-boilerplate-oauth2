package auth

import "time"

type AuthProvider struct {
	ProviderName   string    `json:"provider_name,omitempty" db:"provider_name"`
	ProviderUserID string    `json:"provider_user_id,omitempty" db:"provider_user_id"`
	AccessToken    string    `json:"access_token,omitempty" db:"access_token"`
	RefreshToken   string    `json:"refresh_token,omitempty" db:"refresh_token"`
	Expiry         time.Time `json:"expiry,omitempty" db:"expiry"`
	UserID         string    `json:"user_id,omitempty" db:"user_id"`
	CreatedAt      time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Return the provider name
func (a *AuthProvider) GetName() string {
	return a.ProviderName
}

// Quick check to see if AccessToken is expired
func (a *AuthProvider) IsExpired() bool {
	return a.Expiry.Before(time.Now())
}

// Checks if the linked profile has been updated
func (a *AuthProvider) IsUpdated() bool {
	return a.UpdatedAt.After(a.CreatedAt)
}
