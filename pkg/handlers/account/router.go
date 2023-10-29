package account

import (
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/database"
)

func RegisterRouter(apiRouter *echo.Group) {
	controller := newAccountController(
		newAccountService(database.NewTransactionExecutor(), newAccountRepository()),
	)

	accountRouter := apiRouter.Group("/account")
	accountRouter.POST("/sign-up", controller.signUp)
	accountRouter.POST("/sign-in", controller.signIn)
}
