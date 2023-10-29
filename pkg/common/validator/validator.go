package validator

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Validate(data interface{}) *echo.HTTPError {
	validate := validator.New()
	errs := validate.Struct(data)
	if errs != nil {
		errResult := make(map[string]string)
		for _, err := range errs.(validator.ValidationErrors) {
			errResult[err.Field()] = fmt.Sprintf("ValidationCondition: %s, Got: %v", err.ActualTag(), err.Value())
		}

		return echo.NewHTTPError(http.StatusBadRequest, errResult)
	}

	return nil
}
