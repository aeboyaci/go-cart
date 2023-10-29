package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-cart/pkg/common/jwt_service"
	"go-cart/pkg/common/types"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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

	token := jwt_service.SignJwt("testing@gocart.app", types.Customer, time.Now().Add(24*time.Hour))
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

	token := jwt_service.SignJwt("testing@gocart.app", types.Customer, time.Now().Add(24*time.Hour))
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

	token := jwt_service.SignJwt("testing@gocart.app", types.Admin, time.Now().Add(24*time.Hour))
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
