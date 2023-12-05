package application

import (
	"context"
	"fmt"
	"time"

	"github.com/dany0814/go-hexagonal/internal/core/application/dto"
	"github.com/dany0814/go-hexagonal/internal/core/domain"
	outdb "github.com/dany0814/go-hexagonal/internal/core/ports/driven"
	"github.com/dany0814/go-hexagonal/pkg/encryption"
	"github.com/dany0814/go-hexagonal/pkg/uidgen"
)

type UserService struct {
	userDB outdb.UserDB
}

func NewUserService(userDB outdb.UserDB) UserService {
	return UserService{
		userDB: userDB,
	}
}

func (usrv UserService) Register(ctx context.Context, user dto.User) (*dto.User, error) {
	id := uidgen.New().New()

	newuser, err := domain.NewUser(id, user.Name, user.Lastname, user.Email, user.Password)

	if err != nil {
		return nil, err
	}

	pass, err := encryption.HashAndSalt(user.Password)

	if err != nil {
		return nil, err
	}

	passencrypted, _ := domain.NewUserPassword(pass)

	newuser.Password = passencrypted
	newuser.CreatedAt = time.Now()
	newuser.UpdatedAt = time.Now()

	fmt.Println("user app core: ", newuser)
	err = usrv.userDB.Create(ctx, newuser)

	if err != nil {
		return nil, err
	}

	user.ID = id
	return &user, nil
}
