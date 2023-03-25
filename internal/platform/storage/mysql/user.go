package mysqldb

import "time"

const (
	sqlUserTable = "users"
)

type SqlUser struct {
	ID        string     `db:"user_id"`
	Name      string     `db:"name"`
	Lastname  string     `db:"lastname"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	Dni       string     `db:"dni"`
	Phone     string     `db:"phone"`
	State     string     `db:"state"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
