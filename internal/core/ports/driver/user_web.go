package driverport

import (
	_ "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

type UserAPI interface {
	SignInHandler() echo.HandlerFunc
	GetAllUserHandler() echo.HandlerFunc
}
