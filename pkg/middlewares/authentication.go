package middlewares

import (
	"github.com/labstack/echo/v4"
	jwtService "go-cart/pkg/common/jwt_service"
	"go-cart/pkg/common/types"
	"net/http"
)

func EnforceAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			headerValue := ctx.Request().Header.Get("Authorization")
			claims, err := jwtService.GetTokenClaims(headerValue)
			if err != nil {
				return err
			}

			ctx.Set("email", claims["email"])
			return next(ctx)
		}
	}
}

func EnforceAdminAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			headerValue := ctx.Request().Header.Get("Authorization")
			claims, err := jwtService.GetTokenClaims(headerValue)
			if err != nil {
				return err
			}
			role := claims["role"]
			if role != string(types.Admin) {
				return echo.NewHTTPError(http.StatusUnauthorized, "you are not allowed to do this operation")
			}

			ctx.Set("email", claims["email"])
			return next(ctx)
		}
	}
}
