package env

import (
	"fmt"
	"os"
)

var (
	DB_URL     string
	JWT_SECRET string
)

func Load() error {
	var ok bool
	errorMessage := "cannot load environment variable %s"

	DB_URL, ok = os.LookupEnv("DB_URL")
	if !ok {
		return fmt.Errorf(errorMessage, "DB_URL")
	}

	JWT_SECRET, ok = os.LookupEnv("JWT_SECRET")
	if !ok {
		return fmt.Errorf(errorMessage, "JWT_SECRET")
	}

	return nil
}
