package dto

import "time"

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Lastname  string     `json:"lastname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Dni       string     `json:"dni"`
	Phone     string     `json:"phone,omitempty"`
	State     string     `json:"state,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
