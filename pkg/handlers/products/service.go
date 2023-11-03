package products

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go-cart/pkg/common/database"
	"go-cart/pkg/models"
	"gorm.io/gorm"
	"net/http"
)

type Service interface {
	AddProduct(product models.Product) error
	UpdateProduct(productId string, product models.Product) (models.Product, error)
	DeleteProduct(productId string) error
	GetAllProducts() ([]models.Product, error)
	GetProductById(productId string) (models.Product, error)
}

type productServiceImpl struct {
	txExecutor database.TransactionExecutor
	repository Repository
}

func NewService(txExecutor database.TransactionExecutor, repository Repository) Service {
	return productServiceImpl{
		txExecutor: txExecutor,
		repository: repository,
	}
}

func (s productServiceImpl) AddProduct(product models.Product) error {
	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		err := s.repository.Save(tx, product)
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

func (s productServiceImpl) UpdateProduct(productId string, product models.Product) (models.Product, error) {
	if product.Price < 0 {
		return models.Product{}, echo.NewHTTPError(http.StatusBadRequest, "price has to be a positive number")
	}
	if product.QuantityAvailable < 0 {
		return models.Product{}, echo.NewHTTPError(http.StatusBadRequest, "quantity has to be positive number")
	}

	var result models.Product
	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		_, err := s.repository.FindById(tx, productId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, "product does not exist")
			}
			return err
		}

		result, err = s.repository.UpdateById(tx, productId, product)
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

func (s productServiceImpl) DeleteProduct(productId string) error {
	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		_, err := s.repository.FindById(tx, productId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, "product does not exist")
			}
			return err
		}

		err = s.repository.DeleteById(tx, productId)
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

func (s productServiceImpl) GetAllProducts() ([]models.Product, error) {
	products := make([]models.Product, 0)

	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		var err error
		products, err = s.repository.FindAll(tx)
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

func (s productServiceImpl) GetProductById(productId string) (models.Product, error) {
	var product models.Product

	err := s.txExecutor.Exec(func(tx *gorm.DB) error {
		var err error
		product, err = s.repository.FindById(tx, productId)
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
