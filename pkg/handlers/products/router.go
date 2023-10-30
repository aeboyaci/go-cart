package products

import (
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/database"
	"go-cart/pkg/middlewares"
)

func RegisterRouter(apiRouter *echo.Group) {
	controller := newProductController(
		newProductService(database.NewTransactionExecutor(), newProductRepository()),
	)

	adminProtectedProductsRouter := apiRouter.Group("/products", middlewares.EnforceAdminAuthentication())
	adminProtectedProductsRouter.POST("/", controller.addProduct)
	adminProtectedProductsRouter.PUT("/:id", controller.updateProduct)
	adminProtectedProductsRouter.DELETE("/:id", controller.deleteProduct)

	productsRouter := apiRouter.Group("/products")
	productsRouter.GET("/", controller.getAllProducts)
	productsRouter.GET("/:id", controller.getProductById)
}
