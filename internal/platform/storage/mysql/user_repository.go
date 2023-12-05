package mysqldb

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/dany0814/go-hexagonal/internal/core/domain"
)

type UserRepository struct {
	db        *gorm.DB
	dbTimeout time.Duration
}

// NewUserRepository initializes a MySQL-based implementation of UserRepository.
func NewUserRepository(db *gorm.DB, dbTimeout time.Duration) *UserRepository {
	return &UserRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the adapter userRepository interface.
func (r *UserRepository) Save(ctx context.Context, user domain.User) error {
	fmt.Println("Guardando datos...")
	fmt.Println("Datos para la DB", user)
	// var book Book
	if result := r.db.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}
	return nil
}
