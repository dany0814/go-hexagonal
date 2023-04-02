package drivenport

import (
	"context"

	"github.com/dany0814/go-hexagonal/internal/core/application/dto"
	"github.com/dany0814/go-hexagonal/internal/core/domain"
)

type UserDB interface {
	Create(ctx context.Context, user domain.User) error
	FindAll(ctx context.Context) ([]*dto.User, error)
}
