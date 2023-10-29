package jwt_service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/env"
	"go-cart/pkg/common/types"
	"net/http"
	"strings"
	"time"
)

func GetTokenClaims(headerValue string) (jwt.MapClaims, error) {
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

func SignJwt(email string, role types.UserRole, expireTime time.Time) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expireTime.Unix()
	claims["email"] = email
	claims["role"] = role

	tokenString, err := token.SignedString([]byte(env.JWT_SECRET))
	if err != nil {
		return ""
	}

	return tokenString
}
