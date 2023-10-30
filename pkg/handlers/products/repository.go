package products

import (
	"go-cart/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productRepository interface {
	save(tx *gorm.DB, product models.Product) error
	findById(tx *gorm.DB, productId string) (models.Product, error)
	updateById(tx *gorm.DB, productId string, product models.Product) (models.Product, error)
	deleteById(tx *gorm.DB, productId string) error
	findAll(tx *gorm.DB) ([]models.Product, error)
}

type productRepositoryImpl struct {
}

func newProductRepository() productRepository {
	return productRepositoryImpl{}
}

func (p productRepositoryImpl) save(tx *gorm.DB, product models.Product) error {
	return tx.Create(&product).Error
}

func (p productRepositoryImpl) findById(tx *gorm.DB, productId string) (models.Product, error) {
	var product models.Product
	err := tx.
		Model(&models.Product{}).
		Where("id = ?", productId).
		Take(&product).
		Error
	return product, err
}

func (p productRepositoryImpl) updateById(tx *gorm.DB, productId string, product models.Product) (models.Product, error) {
	err := tx.
		Model(&models.Product{}).
		Clauses(clause.Returning{}).
		Where("id = ?", productId).
		Updates(&product).
		Error
	return product, err
}

func (p productRepositoryImpl) deleteById(tx *gorm.DB, productId string) error {
	return tx.
		Delete(&models.Product{}, productId).
		Error
}

func (p productRepositoryImpl) findAll(tx *gorm.DB) ([]models.Product, error) {
	var products []models.Product
	err := tx.
		Model(&models.Product{}).
		Find(&products).
		Error
	return products, err
}
