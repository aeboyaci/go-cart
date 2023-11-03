package products

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/database"
	"go-cart/pkg/models"
	"gorm.io/gorm"
	"net/http"
)

type productService interface {
	addProduct(product models.Product) error
	updateProduct(productId string, product models.Product) (models.Product, error)
	deleteProduct(productId string) error
	getAllProducts() ([]models.Product, error)
	getProductById(productId string) (models.Product, error)
}

type productServiceImpl struct {
	txExecutor database.TransactionExecutor
	repository productRepository
}

func newProductService(txExecutor database.TransactionExecutor, repository productRepository) productService {
	return productServiceImpl{
		txExecutor: txExecutor,
		repository: repository,
	}
}

func (s productServiceImpl) addProduct(product models.Product) error {
	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		err := s.repository.save(tx, product)
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

func (s productServiceImpl) updateProduct(productId string, product models.Product) (models.Product, error) {
	if product.Price <= 0 {
		return models.Product{}, echo.NewHTTPError(http.StatusBadRequest, "price has to be a positive number")
	}
	if product.QuantityAvailable <= 0 {
		return models.Product{}, echo.NewHTTPError(http.StatusBadRequest, "quantity has to be positive number")
	}

	var result models.Product
	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		_, err := s.repository.findById(tx, productId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, "product does not exist")
			}
			return err
		}

		result, err = s.repository.updateById(tx, productId, product)
		if err != nil {
			return err
		}

		return nil
	}, false)
	if err != nil {
		return models.Product{}, err
	}

	return result, nil
}

func (s productServiceImpl) deleteProduct(productId string) error {
	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		_, err := s.repository.findById(tx, productId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, "product does not exist")
			}
			return err
		}

		err = s.repository.deleteById(tx, productId)
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

func (s productServiceImpl) getAllProducts() ([]models.Product, error) {
	products := make([]models.Product, 0)

	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		var err error
		products, err = s.repository.findAll(tx)
		if err != nil {
			return err
		}

		return nil
	}, true)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s productServiceImpl) getProductById(productId string) (models.Product, error) {
	var product models.Product

	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		var err error
		product, err = s.repository.findById(tx, productId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, "product does not exist")
			}
			return err
		}

		return nil
	}, true)
	if err != nil {
		return product, err
	}

	return product, nil
}
