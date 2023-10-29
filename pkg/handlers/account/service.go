package account

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/database"
	"go-cart/pkg/common/types"
	"go-cart/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"net/mail"
)

type accountService interface {
	signUp(user models.User) (echo.Map, error)
}

type accountServiceImpl struct {
	txExecutor database.TransactionExecutor
	repository accountRepository
}

func newAccountService(txExecutor database.TransactionExecutor, repository accountRepository) accountService {
	return accountServiceImpl{
		txExecutor: txExecutor,
		repository: repository,
	}
}

func (s accountServiceImpl) signUp(user models.User) (echo.Map, error) {
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return echo.Map{}, echo.NewHTTPError(http.StatusBadRequest, "invalid email given")
	}

	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		err := s.repository.existsByEmail(tx, user.Email)
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
		err = s.repository.save(tx, user)
		if err != nil {
			return err
		}

		return nil
	}, false)
	return echo.Map{"success": true, "data": "user created successfully"}, err
}
