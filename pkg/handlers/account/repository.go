package account

import (
	"go-cart/pkg/models"
	"gorm.io/gorm"
)

type accountRepository interface {
	existsByEmail(tx *gorm.DB, email string) error
	save(tx *gorm.DB, user models.User) error
}

type accountRepositoryImpl struct {
}

func newAccountRepository() accountRepository {
	return accountRepositoryImpl{}
}

func (r accountRepositoryImpl) existsByEmail(tx *gorm.DB, email string) error {
	var result models.User
	err := tx.Model(&models.User{}).Where("email = ?", email).Take(&result).Error
	return err
}

func (r accountRepositoryImpl) save(tx *gorm.DB, user models.User) error {
	return tx.Create(&user).Error
}
