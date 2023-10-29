package jwt_service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-cart/pkg/common/types"
	"testing"
	"time"
)

func Test_GetTokenClaims_ValidToken(t *testing.T) {
	token := SignJwt("testing@gocart.app", types.Customer, time.Now().Add(24*time.Hour))
	assert.NotEqual(t, "", token)

	claims, err := GetTokenClaims(fmt.Sprintf("Bearer %s", token))
	assert.Nil(t, err)
	assert.Equal(t, claims["email"], "testing@gocart.app")
	assert.Equal(t, claims["role"], types.Customer.String())
}

func Test_GetTokenClaims_ExpiredToken(t *testing.T) {
	token := SignJwt("testing@gocart.app", types.Customer, time.Now().Add(-24*time.Hour))
	assert.NotEqual(t, "", token)

	claims, err := GetTokenClaims(fmt.Sprintf("Bearer %s", token))
	assert.NotNil(t, err)
	assert.Nil(t, claims["email"])
	assert.Nil(t, claims["role"])
}

func Test_GetTokenClaims_InvalidHeader(t *testing.T) {
	claims, err := GetTokenClaims("random value")
	assert.NotNil(t, err)
	assert.Nil(t, claims["email"])
	assert.Nil(t, claims["role"])
}

func Test_GetTokenClaims_EmptyHeader(t *testing.T) {
	claims, err := GetTokenClaims("")
	assert.NotNil(t, err)
	assert.Nil(t, claims["email"])
	assert.Nil(t, claims["role"])
}
