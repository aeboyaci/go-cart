package account

import (
	"go-cart/pkg/models"
	"gorm.io/gorm"
)

type Repository interface {
	FindByEmail(tx *gorm.DB, email string) (models.User, error)
	Save(tx *gorm.DB, user models.User) error
}

type accountRepositoryImpl struct {
}

func NewRepository() Repository {
	return accountRepositoryImpl{}
}

func (r accountRepositoryImpl) FindByEmail(tx *gorm.DB, email string) (models.User, error) {
	var result models.User
	err := tx.Model(&models.User{}).Where("email = ?", email).Take(&result).Error
	return result, err
}

func (r accountRepositoryImpl) Save(tx *gorm.DB, user models.User) error {
	return tx.Create(&user).Error
}
