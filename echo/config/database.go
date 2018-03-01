package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func OpenDB() *gorm.DB {

	dbHost := "database"
	dbPort := "5432"
	dbName := "workshop"
	dbUser := "postgres"
	dbPass := "pass"

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	driver := "postgres"

	db, err := gorm.Open(driver, connectionString)
	if err != nil {
		panic("failed to connect database " + connectionString)
	}

	return db
}
