package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-cart/pkg/common/database"
	"go-cart/pkg/common/fixtures"
	"go-cart/pkg/handlers/products"
	"gorm.io/gorm"
	"testing"
)

var (
	fixtureFolder string                       = "./fixtures/"
	txExecutor    database.TransactionExecutor = database.NewTransactionExecutor()
	underTest     products.Repository          = products.NewRepository()
)

func Test_FindById_NotFound(t *testing.T) {
	defer fixtures.ClearTables("products")
	fixtures.LoadFixtures(t, fixtureFolder, "products.yml")

	err := txExecutor.Exec(func(tx *gorm.DB) error {
		_, err := underTest.FindById(tx, "5c6f4642-ff78-4eda-b0a9-647843e7eabd")
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		assert.NotNil(t, err)
		assert.Equal(t, err, gorm.ErrRecordNotFound)
		return nil
	}, true)

	assert.Nil(t, err)
}

func Test_FindById_Found(t *testing.T) {
	defer fixtures.ClearTables("products")
	fixtures.LoadFixtures(t, fixtureFolder, "products.yml")

	err := txExecutor.Exec(func(tx *gorm.DB) error {
		product, err := underTest.FindById(tx, "9020ac0b-2a5a-4f40-8211-1f8871a93fb1")
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		assert.Nil(t, err)
		assert.Equal(t, product.Name, "Product - 1")
		return nil
	}, true)

	assert.Nil(t, err)
}
