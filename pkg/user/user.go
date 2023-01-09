package user

import (
	"time"
)

type User struct {
	ID        string    `json:"id,omtiempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	IsActive  bool      `json:"is_active,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *User) IsUpdated() bool {
	if u.CreatedAt == u.UpdatedAt {
		return false
	}
	return true
}
