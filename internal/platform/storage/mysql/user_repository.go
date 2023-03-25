package mysqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dany0814/go-hexagonal/internal/core/application/dto"
	"github.com/huandu/go-sqlbuilder"
)

type UserRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewCourseRepository initializes a MySQL-based implementation of mooc.CourseRepository.
func NewUserRepository(db *sql.DB, dbTimeout time.Duration) *UserRepository {
	return &UserRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the adapter userRepository interface.
func (r *UserRepository) Save(ctx context.Context, user SqlUser) error {
	userSQLStruct := sqlbuilder.NewStruct(new(SqlUser))
	query, args := userSQLStruct.InsertInto(sqlUserTable, SqlUser{
		ID:        user.ID,
		Name:      user.Name,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
		Dni:       user.Dni,
		State:     user.State,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Error trying to persist course on database: %v", err)
	}

	return nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*dto.User, error) {
	var allUser []*dto.User

	userSQLStruct := sqlbuilder.NewStruct(new(SqlUser))
	query := userSQLStruct.SelectFrom(sqlUserTable)
	results, err := r.db.Query(query.String())

	if err != nil {
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var sqlUser SqlUser
		if err := results.Scan(
			&sqlUser.ID,
			&sqlUser.Name,
			&sqlUser.Lastname,
			&sqlUser.Email,
			&sqlUser.Password,
			&sqlUser.Dni,
			&sqlUser.Phone,
			&sqlUser.State,
			&sqlUser.CreatedAt,
			&sqlUser.UpdatedAt,
			&sqlUser.DeletedAt,
		); err != nil {
			return nil, err
		}
		sqlUser.Password = ""
		user := dto.User(sqlUser)
		allUser = append(allUser, &user)
	}
	return allUser, nil
}
