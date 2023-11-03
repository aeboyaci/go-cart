package account

import (
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/database"
)

func RegisterRouter(apiRouter *echo.Group) {
	controller := NewController(
		NewService(database.NewTransactionExecutor(), NewRepository()),
	)

	accountRouter := apiRouter.Group("/account")
	accountRouter.POST("/sign-up", controller.SignUp)
	accountRouter.POST("/sign-in", controller.SignIn)
}
