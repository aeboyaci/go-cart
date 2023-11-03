package account

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/database"
	jwtService "go-cart/pkg/common/jwt_service"
	"go-cart/pkg/common/types"
	"go-cart/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Service interface {
	SignUp(user models.User) error
	SignIn(userParams SignInParams) (string, error)
}

type accountServiceImpl struct {
	txExecutor database.TransactionExecutor
	repository Repository
}

func NewService(txExecutor database.TransactionExecutor, repository Repository) Service {
	return accountServiceImpl{
		txExecutor: txExecutor,
		repository: repository,
	}
}

func (s accountServiceImpl) SignUp(user models.User) error {
	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		_, err := s.repository.FindByEmail(tx, user.Email)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		emailExists := !errors.Is(err, gorm.ErrRecordNotFound)
		if emailExists {
			return echo.NewHTTPError(http.StatusBadRequest, "email taken")
		}

		user.Role = types.Customer
		hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "cannot hash password")
		}
		user.Password = string(hashedPasswordBytes)
		err = s.repository.Save(tx, user)
		if err != nil {
			return err
		}

		return nil
	}, false)
	if err != nil {
		return err
	}

	return nil
}

func (s accountServiceImpl) SignIn(userParams SignInParams) (string, error) {
	var token string
	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		dbUser, err := s.repository.FindByEmail(tx, userParams.Email)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, "email or password is incorrect")
			}
			return err
		}
		err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(userParams.Password))
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "email or password is incorrect")
		}

		token = jwtService.SignJwt(dbUser.Email, dbUser.Role, time.Now().Add(24*time.Hour))
		return nil
	}, true)
	if err != nil {
		return token, err
	}

	return token, nil
}
