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
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON body sent")
	}
	if err := validator.Validate(user); err != nil {
		return err
	}

	err := c.service.signUp(user)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, echo.Map{"success": true, "data": "user created successfully"})
}

func (c accountController) signIn(ctx echo.Context) error {
	var userParams signInParams
	if err := ctx.Bind(&userParams); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON body sent")
	}
	if err := validator.Validate(userParams); err != nil {
		return err
	}

	token, err := c.service.signIn(userParams)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "data": token})
}
