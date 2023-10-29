package account

import (
	"go-cart/pkg/models"
	"gorm.io/gorm"
)

type accountRepository interface {
	findByEmail(tx *gorm.DB, email string) (models.User, error)
	save(tx *gorm.DB, user models.User) error
}

type accountRepositoryImpl struct {
}

func newAccountRepository() accountRepository {
	return accountRepositoryImpl{}
}

func (r accountRepositoryImpl) findByEmail(tx *gorm.DB, email string) (models.User, error) {
	var result models.User
	err := tx.Model(&models.User{}).Where("email = ?", email).Take(&result).Error
	return result, err
}

func (r accountRepositoryImpl) save(tx *gorm.DB, user models.User) error {
	return tx.Create(&user).Error
}
