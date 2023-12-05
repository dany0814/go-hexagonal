package mysqldb

import "time"

type Book struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Lastname  string     `json:"lastname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
