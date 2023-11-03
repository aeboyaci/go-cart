// Code generated by mockery v2.36.0. DO NOT EDIT.

package tests

import (
	models "go-cart/pkg/models"

	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// mockRepositoryImpl is an autogenerated mock type for the Repository type
type mockRepositoryImpl struct {
	mock.Mock
}

// DeleteById provides a mock function with given fields: tx, productId
func (_m *mockRepositoryImpl) DeleteById(tx *gorm.DB, productId string) error {
	ret := _m.Called(tx, productId)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) error); ok {
		r0 = rf(tx, productId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: tx
func (_m *mockRepositoryImpl) FindAll(tx *gorm.DB) ([]models.Product, error) {
	ret := _m.Called(tx)

	var r0 []models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB) ([]models.Product, error)); ok {
		return rf(tx)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB) []models.Product); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: tx, productId
func (_m *mockRepositoryImpl) FindById(tx *gorm.DB, productId string) (models.Product, error) {
	ret := _m.Called(tx, productId)

	var r0 models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) (models.Product, error)); ok {
		return rf(tx, productId)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) models.Product); ok {
		r0 = rf(tx, productId)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string) error); ok {
		r1 = rf(tx, productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: tx, product
func (_m *mockRepositoryImpl) Save(tx *gorm.DB, product models.Product) error {
	ret := _m.Called(tx, product)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, models.Product) error); ok {
		r0 = rf(tx, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateById provides a mock function with given fields: tx, productId, product
func (_m *mockRepositoryImpl) UpdateById(tx *gorm.DB, productId string, product models.Product) (models.Product, error) {
	ret := _m.Called(tx, productId, product)

	var r0 models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, models.Product) (models.Product, error)); ok {
		return rf(tx, productId, product)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, models.Product) models.Product); ok {
		r0 = rf(tx, productId, product)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string, models.Product) error); ok {
		r1 = rf(tx, productId, product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockRepositoryImpl creates a new instance of mockRepositoryImpl. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockRepositoryImpl(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockRepositoryImpl {
	mock := &mockRepositoryImpl{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}