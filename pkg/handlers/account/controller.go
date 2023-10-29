package account

import (
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/validator"
	"go-cart/pkg/models"
	"net/http"
)

type accountController struct {
	service accountService
}

func newAccountController(accountService accountService) accountController {
	return accountController{
		service: accountService,
	}
}

func (c accountController) signUp(ctx echo.Context) error {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return err
	}
	if err := validator.Validate(user); err != nil {
		return err
	}

	result, err := c.service.signUp(user)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, result)
}

func (c accountController) signIn(ctx echo.Context) error {
	var userParams signInParams
	if err := ctx.Bind(&userParams); err != nil {
		return err
	}
	if err := validator.Validate(userParams); err != nil {
		return err
	}

	result, err := c.service.signIn(userParams)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}
