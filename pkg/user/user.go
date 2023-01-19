package user

import (
	"time"
)

type User struct {
	ID        string    `json:"id,omtiempty" db:"id"`
	Name      string    `json:"name,omitempty" db:"name"`
	Username  string    `json:"username,omitempty" db:"username"`
	Email     string    `json:"email,omitempty" db:"email"`
	IsActive  bool      `json:"is_active,omitempty" db:"is_active"`
	Password  string    `json:"password,omitempty" db:"hashed_password"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (u *User) IsUpdated() bool {
	if u.CreatedAt == u.UpdatedAt {
		return false
	}
	return true
}
