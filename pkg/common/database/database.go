package database

import (
	"fmt"
	"go-cart/pkg/common/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Initialize() error {
	var err error
	db, err = gorm.Open(postgres.Open(env.DB_URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("cannot connect to the database. Error: %s", err.Error())
	}

	return autoMigrate()
}

func getClient() *gorm.DB {
	if db == nil {
		panic("database connection has not been established yet! please connect to the database first")
	}

	return db
}

func autoMigrate() error {
	return nil
}
