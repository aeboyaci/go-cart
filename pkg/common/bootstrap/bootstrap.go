package bootstrap

import (
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/database"
	"go-cart/pkg/common/env"
	"go-cart/pkg/handlers/account"
	"go-cart/pkg/handlers/products"
	"net/http"
)

func Initialize() error {
	err := env.Load()
	if err != nil {
		return err
	}

	err = database.Initialize()
	if err != nil {
		return err
	}

	return nil
}

func RegisterRouters(e *echo.Echo) {
	e.HTTPErrorHandler = customHttpErrorHandler
	apiRouter := e.Group("/api")

	account.RegisterRouter(apiRouter)
	products.RegisterRouter(apiRouter)
}

func customHttpErrorHandler(err error, c echo.Context) {
	statusCode := http.StatusInternalServerError
	var errorMessage interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		statusCode = he.Code
		errorMessage = he.Message
	}
	c.Logger().Error(err)

	err = c.JSON(statusCode, echo.Map{
		"success": false,
		"error":   errorMessage,
	})
	if err != nil {
		c.Logger().Error(err)
	}
}
