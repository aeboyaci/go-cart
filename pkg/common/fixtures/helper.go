package fixtures

import (
	"fmt"
	"github.com/go-testfixtures/testfixtures"
	"go-cart/pkg/common/database"
	"go-cart/pkg/common/env"
	"log"
	"testing"
)

func init() {
	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}
	err = database.Initialize()
	if err != nil {
		log.Fatal(err)
	}
}

func LoadFixtures(t *testing.T, fixtureFolder string, fileNames ...string) {
	helper := testfixtures.PostgreSQL{
		UseAlterConstraint: false,
		SkipResetSequences: true,
	}

	for i := range fileNames {
		fileNames[i] = fixtureFolder + fileNames[i]
	}

	db := database.GetClient()
	dbClient, _ := db.DB()
	context, err := testfixtures.NewFiles(dbClient, &helper, fileNames...)

	if err != nil {
		fmt.Println("test fixtures context could not be created", err)
		t.Fail()
		return
	} else {
		if err := context.Load(); err != nil {
			fmt.Println("test fixtures could not be loaded", err)
			t.Fail()
			return
		}
	}
}

func ClearTables(tableNames ...string) {
	db := database.GetClient()
	for i := range tableNames {
		err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", tableNames[i])).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}
