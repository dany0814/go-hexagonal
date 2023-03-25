package repository

import (
	"context"

	"github.com/dany0814/go-hexagonal/internal/core/application/dto"
	mysqldb "github.com/dany0814/go-hexagonal/internal/platform/storage/mysql"
)

type UserRepository interface {
	Save(ctx context.Context, sqluser mysqldb.SqlUser) error
	FindAll(ctx context.Context) ([]*dto.User, error)
}
