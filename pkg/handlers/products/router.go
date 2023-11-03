package products

import (
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/database"
	"go-cart/pkg/middlewares"
)

func RegisterRouter(apiRouter *echo.Group) {
	controller := NewController(
		NewService(database.NewTransactionExecutor(), newProductRepository()),
	)

	adminProtectedProductsRouter := apiRouter.Group("/products", middlewares.EnforceAdminAuthentication())
	adminProtectedProductsRouter.POST("/", controller.AddProduct)
	adminProtectedProductsRouter.PUT("/:id", controller.UpdateProduct)
	adminProtectedProductsRouter.DELETE("/:id", controller.DeleteProduct)

	productsRouter := apiRouter.Group("/products")
	productsRouter.GET("/", controller.GetAllProducts)
	productsRouter.GET("/:id", controller.GetProductById)
}
