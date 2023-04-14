package domain

import (
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/dany0814/go-hexagonal/pkg/uidgen"
)

var ErrUserConflict = errors.New("user already exists")
var ErrInvalidUserID = errors.New("invalid User ID")
var ErrInvalidUserEmail = errors.New("invalid Email")
var ErrInvalidUserPassword = errors.New("invalid Password")
var ErrEmptyName = errors.New("the field name is required")

// NewUserID function to instantiate the initial value for UserID

type UserID struct {
	value string
}

func NewUserID(value string) (UserID, error) {
	v, err := uidgen.Parse(value)
	if err != nil {
		return UserID{}, fmt.Errorf("%w: %s", ErrInvalidUserID, value)
	}
	return UserID{
		value: v,
	}, nil
}

func (id UserID) String() string {
	return id.value
}

// NewUserEmail function to instantiate the initial value for UserEmail

type UserEmail struct {
	value string
}

func NewUserEmail(value string) (UserEmail, error) {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return UserEmail{}, fmt.Errorf("%w: %s", ErrInvalidUserEmail, value)
	}
	return UserEmail{
		value: value,
	}, nil
}

func (email UserEmail) String() string {
	return email.value
}

// NewUserPassword function to instantiate the initial value for UserPassword

type UserPassword struct {
	value string
}

func NewUserPassword(value string) (UserPassword, error) {
	if value == "" {
		return UserPassword{}, fmt.Errorf("%w: %s", ErrInvalidUserPassword, value)
	}
	return UserPassword{
		value: value,
	}, nil
}

func (pass UserPassword) String() string {
	return pass.value
}

// NewUserUsername function to instantiate the initial value for UserUsername

// NewUser function to instantiate the initial value for User

type User struct {
	ID        UserID
	Name      string
	Lastname  string
	Email     UserEmail
	Password  UserPassword
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewUser(userID, name, lastname, email, password string) (User, error) {
	idVo, err := NewUserID(userID)
	if err != nil {
		return User{}, err
	}

	if name == "" {
		return User{}, fmt.Errorf("%w: %s", ErrEmptyName, name)
	}

	emailVo, err := NewUserEmail(email)
	if err != nil {
		return User{}, err
	}

	passwordVo, err := NewUserPassword(password)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:       idVo,
		Name:     name,
		Lastname: lastname,
		Email:    emailVo,
		Password: passwordVo,
	}, nil
}

func (u User) UserID() UserID {
	return u.UserID()
}
