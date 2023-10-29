package database

import "gorm.io/gorm"

type mockTransactionExecutorImpl struct {
}

func NewMockTransactionExecutor() TransactionExecutor {
	return mockTransactionExecutorImpl{}
}

func (t mockTransactionExecutorImpl) Exec(fn func(tx *gorm.DB) error, readOnly bool) error {
	return fn(nil)
}
