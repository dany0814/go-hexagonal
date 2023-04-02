package drivenport

import (
	"context"

	"github.com/dany0814/go-hexagonal/internal/core/domain"
)

type UserDB interface {
	Create(ctx context.Context, user domain.User) error
}
