package bootstrap

import (
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/database"
	"go-cart/pkg/common/env"
	"go-cart/pkg/handlers/account"
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

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "PONG!")
	})

	account.RegisterRouter(apiRouter)
}

func customHttpErrorHandler(err error, c echo.Context) {
	statusCode := http.StatusInternalServerError
	var errorMessage interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		statusCode = he.Code
		errorMessage = he.Message
	}

	err = c.JSON(statusCode, echo.Map{
		"success": false,
		"error":   errorMessage,
	})
	if err != nil {
		c.Logger().Error(err)
	}
}
