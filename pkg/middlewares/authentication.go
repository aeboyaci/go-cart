package middlewares

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/env"
	"go-cart/pkg/common/types"
	"net/http"
	"strings"
)

func getTokenClaims(headerValue string) (jwt.MapClaims, error) {
	tokenString := strings.ReplaceAll(headerValue, "Bearer ", "")

	if tokenString == "" {
		return jwt.MapClaims{}, echo.NewHTTPError(http.StatusUnauthorized, "you are not allowed to do this operation")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.JWT_SECRET), nil
	})
	if errors.Is(err, jwt.ErrTokenMalformed) {
		return jwt.MapClaims{}, echo.NewHTTPError(http.StatusUnauthorized, "you are not allowed to do this operation")
	}
	if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return jwt.MapClaims{}, echo.NewHTTPError(http.StatusUnauthorized, "your authentication token has expired")
	}

	return token.Claims.(jwt.MapClaims), nil
}

func EnforceAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			headerValue := ctx.Request().Header.Get("Authorization")
			claims, err := getTokenClaims(headerValue)
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
			claims, err := getTokenClaims(headerValue)
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
