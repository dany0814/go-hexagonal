package repository

import (
	"context"

	mysqldb "github.com/dany0814/go-hexagonal/internal/platform/storage/mysql"
)

type UserRepository interface {
	Save(ctx context.Context, sqluser mysqldb.SqlUser) error
}
