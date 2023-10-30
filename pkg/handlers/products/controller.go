package products

import (
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/validator"
	"go-cart/pkg/models"
	"net/http"
)

type productController struct {
	service productService
}

func newProductController(productService productService) productController {
	return productController{
		service: productService,
	}
}

func (c productController) addProduct(ctx echo.Context) error {
	var product models.Product
	if err := ctx.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON body sent")
	}
	if err := validator.Validate(product); err != nil {
		return err
	}

	err := c.service.addProduct(product)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, echo.Map{"success": true, "data": "product added successfully"})
}

func (c productController) updateProduct(ctx echo.Context) error {
	var product models.Product
	if err := ctx.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON body sent")
	}
	// Note that to allow empty values on update, parameter validation is not applied
	productId := ctx.Param("id")

	result, err := c.service.updateProduct(productId, product)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "data": result})
}

func (c productController) deleteProduct(ctx echo.Context) error {
	productId := ctx.Param("id")

	err := c.service.deleteProduct(productId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "data": "product deleted successfully"})
}

func (c productController) getAllProducts(ctx echo.Context) error {
	result, err := c.service.getAllProducts()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "data": result})
}

func (c productController) getProductById(ctx echo.Context) error {
	productId := ctx.Param("id")

	result, err := c.service.getProductById(productId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "data": result})
}
