package account

import (
	"github.com/stretchr/testify/assert"
	"go-cart/pkg/common/database"
	"go-cart/pkg/common/types"
	"go-cart/pkg/models"
	"gorm.io/gorm"
	"testing"
)

var tx *gorm.DB

func setUpSuite(t *testing.T) (accountService, *mockAccountRepositoryImpl) {
	mockRepository := newMockAccountRepositoryImpl(t)
	underTest := newAccountService(
		database.NewMockTransactionExecutor(),
		mockRepository,
	)

	return underTest, mockRepository
}

func Test_SignUp_Fail_InvalidEmailGiven(t *testing.T) {
	underTest, _ := setUpSuite(t)

	user := models.User{
		BaseModel: models.BaseModel{},
		Email:     "invalid email address",
		Role:      types.Customer,
	}
	_, err := underTest.signUp(user)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid email")
}

func Test_SignUp_Fail_EmailTaken(t *testing.T) {
	underTest, mockRepository := setUpSuite(t)

	user := models.User{
		BaseModel: models.BaseModel{},
		Email:     "testing+01@gocart.app",
		Role:      types.Customer,
	}
	mockRepository.
		On("existsByEmail", tx, user.Email).
		Return(nil)

	_, err := underTest.signUp(user)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "email taken")
}

func Test_SignUp_Success(t *testing.T) {
	underTest, mockRepository := setUpSuite(t)

	user := models.User{
		BaseModel: models.BaseModel{},
		Email:     "testing+01@gocart.app",
		Role:      types.Customer,
	}
	mockRepository.
		On("existsByEmail", tx, user.Email).
		Return(gorm.ErrRecordNotFound)
	mockRepository.
		On("save", tx, user).
		Return(nil)

	_, err := underTest.signUp(user)
	assert.Nil(t, err)
}
