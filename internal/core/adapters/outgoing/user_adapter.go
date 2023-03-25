package database

import (
	"context"

	"github.com/dany0814/go-hexagonal/internal/core/adapters/outgoing/repository"
	"github.com/dany0814/go-hexagonal/internal/core/application/dto"
	"github.com/dany0814/go-hexagonal/internal/core/domain"
	mysqldb "github.com/dany0814/go-hexagonal/internal/platform/storage/mysql"
)

type UserAdapter struct {
	userRepository repository.UserRepository
}

func NewUserAdapter(usrep repository.UserRepository) UserAdapter {
	return UserAdapter{
		userRepository: usrep,
	}
}

func (uadpt UserAdapter) Create(ctx context.Context, user domain.User) error {
	sqlUser := mysqldb.SqlUser{
		ID:        user.ID.String(),
		Name:      user.Name,
		Lastname:  user.Lastname,
		Email:     user.Email.String(),
		Password:  user.Password.String(),
		Dni:       user.Dni.String(),
		Phone:     user.Phone,
		State:     user.State,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
	return uadpt.userRepository.Save(ctx, sqlUser)
}

func (uadpt UserAdapter) FindAll(ctx context.Context) ([]*dto.User, error) {
	listUser, err := uadpt.userRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return listUser, nil
}
