package database

import (
	"database/sql"
	"gorm.io/gorm"
)

type TransactionExecutor interface {
	Exec(fn func(tx *gorm.DB) error, readOnly bool) error
}

type transactionExecutorImpl struct {
}

func NewTransactionExecutor() TransactionExecutor {
	return transactionExecutorImpl{}
}

func (t transactionExecutorImpl) Exec(fn func(tx *gorm.DB) error, readOnly bool) error {
	transaction := getClient().Begin(&sql.TxOptions{ReadOnly: readOnly})

	if err := fn(transaction); err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
