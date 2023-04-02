package repository

import (
	"context"

	"github.com/dany0814/go-hexagonal/internal/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, sqluser domain.User) error
}
