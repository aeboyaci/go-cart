package products

import (
	"go-cart/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Save(tx *gorm.DB, product models.Product) error
	FindById(tx *gorm.DB, productId string) (models.Product, error)
	UpdateById(tx *gorm.DB, productId string, product models.Product) (models.Product, error)
	DeleteById(tx *gorm.DB, productId string) error
	FindAll(tx *gorm.DB) ([]models.Product, error)
}

type productRepositoryImpl struct {
}

func newProductRepository() Repository {
	return productRepositoryImpl{}
}

func (p productRepositoryImpl) Save(tx *gorm.DB, product models.Product) error {
	return tx.Create(&product).Error
}

func (p productRepositoryImpl) FindById(tx *gorm.DB, productId string) (models.Product, error) {
	var product models.Product
	err := tx.
		Model(&models.Product{}).
		Where("id = ?", productId).
		Take(&product).
		Error
	return product, err
}

func (p productRepositoryImpl) UpdateById(tx *gorm.DB, productId string, product models.Product) (models.Product, error) {
	err := tx.
		Model(&models.Product{}).
		Clauses(clause.Returning{}).
		Where("id = ?", productId).
		Updates(&product).
		Error
	return product, err
}

func (p productRepositoryImpl) DeleteById(tx *gorm.DB, productId string) error {
	return tx.
		Delete(&models.Product{}, productId).
		Error
}

func (p productRepositoryImpl) FindAll(tx *gorm.DB) ([]models.Product, error) {
	var products []models.Product
	err := tx.
		Model(&models.Product{}).
		Find(&products).
		Error
	return products, err
}
