package tests

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-cart/pkg/common/database"
	"go-cart/pkg/handlers/products"
	"go-cart/pkg/models"
	"gorm.io/gorm"
	"testing"
)

var tx *gorm.DB

func setUpSuite(t *testing.T) (products.Service, *mockRepositoryImpl) {
	mockRepository := newMockRepositoryImpl(t)
	underTest := products.NewService(
		database.NewMockTransactionExecutor(),
		mockRepository,
	)

	return underTest, mockRepository
}

func Test_UpdateProduct_Fail_InvalidPrice(t *testing.T) {
	underTest, _ := setUpSuite(t)

	product, err := underTest.UpdateProduct(mock.Anything, models.Product{
		Price: -1,
	})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "price has to be a positive number")
	assert.Equal(t, product, models.Product{})
}

func Test_UpdateProduct_Fail_InvalidQuantity(t *testing.T) {
	underTest, _ := setUpSuite(t)

	product, err := underTest.UpdateProduct(mock.Anything, models.Product{
		QuantityAvailable: -1,
	})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "quantity has to be positive number")
	assert.Equal(t, product, models.Product{})
}

func Test_UpdateProduct_Fail_ProductNotFound(t *testing.T) {
	underTest, mockRepository := setUpSuite(t)

	id := uuid.New()
	mockRepository.
		On("FindById", tx, id.String()).
		Return(models.Product{}, gorm.ErrRecordNotFound)

	product, err := underTest.UpdateProduct(id.String(), models.Product{
		BaseModel: models.BaseModel{
			ID: id,
		},
		Name: "Product that does not exist in stock",
	})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "product does not exist")
	assert.Equal(t, product, models.Product{})
}

func Test_UpdateProduct_Success(t *testing.T) {
	underTest, mockRepository := setUpSuite(t)

	id := uuid.New()
	oldProduct := models.Product{
		BaseModel: models.BaseModel{
			ID: id,
		},
		Name: "Old Name",
	}
	newProduct := models.Product{
		BaseModel: models.BaseModel{
			ID: id,
		},
		Name: "New Name",
	}

	mockRepository.
		On("FindById", tx, id.String()).
		Return(oldProduct, nil)
	mockRepository.
		On("UpdateById", tx, id.String(), newProduct).
		Return(newProduct, nil)

	result, err := underTest.UpdateProduct(id.String(), newProduct)
	assert.Nil(t, err)
	assert.Equal(t, result.Name, newProduct.Name)
}
