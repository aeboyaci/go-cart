package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-cart/pkg/common/database"
	"go-cart/pkg/common/types"
	"go-cart/pkg/handlers/account"
	"go-cart/pkg/models"
	"gorm.io/gorm"
	"testing"
)

var tx *gorm.DB

func setUpSuite(t *testing.T) (account.Service, *mockRepositoryImpl) {
	mockRepository := newMockRepositoryImpl(t)
	underTest := account.NewService(
		database.NewMockTransactionExecutor(),
		mockRepository,
	)

	return underTest, mockRepository
}

func Test_SignUp_Fail_EmailTaken(t *testing.T) {
	underTest, mockRepository := setUpSuite(t)

	user := models.User{
		BaseModel: models.BaseModel{},
		Email:     "testing+01@gocart.app",
		Role:      types.Customer,
	}
	mockRepository.
		On("FindByEmail", tx, mock.Anything).
		Return(user, nil)

	err := underTest.SignUp(user)
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
		On("FindByEmail", tx, user.Email).
		Return(models.User{}, gorm.ErrRecordNotFound)
	mockRepository.
		On("Save", tx, mock.Anything).
		Return(nil)

	err := underTest.SignUp(user)
	assert.Nil(t, err)
}

func Test_SignIn_Fail_WrongEmail(t *testing.T) {
	underTest, mockRepository := setUpSuite(t)

	userParams := account.SignInParams{
		Email:    "testing+02@gocart.app",
		Password: "123456",
	}
	mockRepository.
		On("FindByEmail", tx, userParams.Email).
		Return(models.User{}, gorm.ErrRecordNotFound)

	_, err := underTest.SignIn(userParams)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "email or password is incorrect")
}

func Test_SignIn_Fail_WrongPassword(t *testing.T) {
	underTest, mockRepository := setUpSuite(t)

	userParams := account.SignInParams{
		Email:    "testing+01@gocart.app",
		Password: "random value",
	}
	user := models.User{
		BaseModel: models.BaseModel{},
		Email:     "testing+01@gocart.app",
		Password:  "$2a$10$QRDlGo6M7ReGc3w6PwU4EuDVLFl1LxjNFKOMF8.Ig8iyWU/vRkqAu",
		Role:      types.Customer,
	}

	mockRepository.
		On("FindByEmail", tx, userParams.Email).
		Return(user, nil)

	_, err := underTest.SignIn(userParams)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "email or password is incorrect")
}

func Test_SignIn_Success(t *testing.T) {
	underTest, mockRepository := setUpSuite(t)

	userParams := account.SignInParams{
		Email:    "testing+01@gocart.app",
		Password: "123456",
	}
	user := models.User{
		BaseModel: models.BaseModel{},
		Email:     "testing+01@gocart.app",
		Password:  "$2a$10$QRDlGo6M7ReGc3w6PwU4EuDVLFl1LxjNFKOMF8.Ig8iyWU/vRkqAu",
		Role:      types.Customer,
	}

	mockRepository.
		On("FindByEmail", tx, userParams.Email).
		Return(user, nil)

	result, err := underTest.SignIn(userParams)
	assert.Nil(t, err)
	assert.NotEqual(t, result, "")
}
