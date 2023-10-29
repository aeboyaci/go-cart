package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-cart/pkg/common/bootstrap"
	"log"
)

func main() {
	err := bootstrap.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	bootstrap.RegisterRouters(e)
	e.Logger.Fatal(e.Start(":1323"))
}
