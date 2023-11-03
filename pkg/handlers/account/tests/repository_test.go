package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-cart/pkg/common/database"
	"go-cart/pkg/common/fixtures"
	"go-cart/pkg/handlers/account"
	"gorm.io/gorm"
	"testing"
)

var (
	fixtureFolder string                       = "./fixtures/"
	txExecutor    database.TransactionExecutor = database.NewTransactionExecutor()
	underTest     account.Repository           = account.NewRepository()
)

func Test_FindByEmail_NotFound(t *testing.T) {
	defer fixtures.ClearTables("users")
	fixtures.LoadFixtures(t, fixtureFolder, "users.yml")

	err := txExecutor.Exec(func(tx *gorm.DB) error {
		_, err := underTest.FindByEmail(tx, "testing+02@gocart.app")
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		assert.NotNil(t, err)
		assert.Equal(t, err, gorm.ErrRecordNotFound)
		return nil
	}, true)

	assert.Nil(t, err)
}

func Test_FindByEmail_Found(t *testing.T) {
	defer fixtures.ClearTables("users")
	fixtures.LoadFixtures(t, fixtureFolder, "users.yml")

	err := txExecutor.Exec(func(tx *gorm.DB) error {
		_, err := underTest.FindByEmail(tx, "testing+01@gocart.app")
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		assert.Nil(t, err)
		return nil
	}, true)

	assert.Nil(t, err)
}
