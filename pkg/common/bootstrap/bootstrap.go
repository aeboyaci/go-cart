package bootstrap

import (
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/database"
	"go-cart/pkg/common/env"
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
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "PONG!")
	})
}
