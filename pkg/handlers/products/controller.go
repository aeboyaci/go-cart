package products

import (
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/validator"
	"go-cart/pkg/models"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(productService Service) Controller {
	return Controller{
		service: productService,
	}
}

func (c Controller) AddProduct(ctx echo.Context) error {
	var product models.Product
	if err := ctx.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON body sent")
	}
	if err := validator.Validate(product); err != nil {
		return err
	}

	err := c.service.AddProduct(product)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, echo.Map{"success": true, "data": "product added successfully"})
}

func (c Controller) UpdateProduct(ctx echo.Context) error {
	var product models.Product
	if err := ctx.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON body sent")
	}
	// Note that to allow empty values on update, parameter validation is not applied
	productId := ctx.Param("id")

	result, err := c.service.UpdateProduct(productId, product)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "data": result})
}

func (c Controller) DeleteProduct(ctx echo.Context) error {
	productId := ctx.Param("id")

	err := c.service.DeleteProduct(productId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "data": "product deleted successfully"})
}

func (c Controller) GetAllProducts(ctx echo.Context) error {
	result, err := c.service.GetAllProducts()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "data": result})
}

func (c Controller) GetProductById(ctx echo.Context) error {
	productId := ctx.Param("id")

	result, err := c.service.GetProductById(productId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"success": true, "data": result})
}
