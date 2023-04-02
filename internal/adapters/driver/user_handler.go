package driveradapt

import (
	"errors"
	"net/http"

	"github.com/dany0814/go-hexagonal/internal/core/application"
	"github.com/dany0814/go-hexagonal/internal/core/application/dto"
	"github.com/dany0814/go-hexagonal/internal/core/domain"
	"github.com/dany0814/go-hexagonal/pkg/helpers"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService application.UserService
	Ctx         echo.Context
}

func NewUserHandler(usrv application.UserService) UserHandler {
	return UserHandler{
		userService: usrv,
	}
}

func (usrh UserHandler) SignInHandler() error {
	ctx := usrh.Ctx
	var req dto.User
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	res, err := usrh.userService.Register(ctx.Request().Context(), req)

	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUserConflict):
			ctx.JSON(http.StatusConflict, err.Error())
			return nil
		default:
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return nil
		}
	}
	ctx.JSON(http.StatusCreated, helpers.DataResponse(0, "User created", res))
	return nil
}
