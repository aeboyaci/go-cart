package middlewares

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-cart/pkg/common/env"
	"go-cart/pkg/common/types"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_GetTokenClaims_ValidToken(t *testing.T) {
	token := signJwt("testing@gocart.app", types.Customer, time.Now().Add(24*time.Hour))
	assert.NotEqual(t, "", token)

	claims, err := getTokenClaims(fmt.Sprintf("Bearer %s", token))
	assert.Nil(t, err)
	assert.Equal(t, claims["email"], "testing@gocart.app")
	assert.Equal(t, claims["role"], types.Customer.String())
}

func Test_GetTokenClaims_ExpiredToken(t *testing.T) {
	token := signJwt("testing@gocart.app", types.Customer, time.Now().Add(-24*time.Hour))
	assert.NotEqual(t, "", token)

	claims, err := getTokenClaims(fmt.Sprintf("Bearer %s", token))
	assert.NotNil(t, err)
	assert.Nil(t, claims["email"])
	assert.Nil(t, claims["role"])
}

func Test_GetTokenClaims_InvalidHeader(t *testing.T) {
	claims, err := getTokenClaims("random value")
	assert.NotNil(t, err)
	assert.Nil(t, claims["email"])
	assert.Nil(t, claims["role"])
}

func Test_GetTokenClaims_EmptyHeader(t *testing.T) {
	claims, err := getTokenClaims("")
	assert.NotNil(t, err)
	assert.Nil(t, claims["email"])
	assert.Nil(t, claims["role"])
}

func Test_EnforceAuthentication_WithoutToken(t *testing.T) {
	e := echo.New()
	e.Use(EnforceAuthentication())

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusUnauthorized)
}

func Test_EnforceAuthentication_ValidToken(t *testing.T) {
	e := echo.New()
	e.Use(EnforceAuthentication())

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	token := signJwt("testing@gocart.app", types.Customer, time.Now().Add(24*time.Hour))
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}

func Test_EnforceAdminAuthentication_WithoutToken(t *testing.T) {
	e := echo.New()
	e.Use(EnforceAdminAuthentication())

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusUnauthorized)
}

func Test_EnforceAdminAuthentication_CustomerToken(t *testing.T) {
	e := echo.New()
	e.Use(EnforceAdminAuthentication())

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	token := signJwt("testing@gocart.app", types.Customer, time.Now().Add(24*time.Hour))
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusUnauthorized)
}

func Test_EnforceAdminAuthentication_AdminToken(t *testing.T) {
	e := echo.New()
	e.Use(EnforceAdminAuthentication())

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	token := signJwt("testing@gocart.app", types.Admin, time.Now().Add(24*time.Hour))
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}

func signJwt(email string, role types.UserRole, expireTime time.Time) string {
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
