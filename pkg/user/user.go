package user

import (
	"fmt"
	"time"
)

type User struct {
	ID        string    `json:"id,omtiempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	IsActive  bool      `json:"is_active,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u *User) IsUpdated() bool {
	if u.CreatedAt == u.UpdatedAt {
		return false
	}
	return true
}
