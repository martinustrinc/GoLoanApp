package main

import (
	"GoLoanApp/config"
	"GoLoanApp/router"
	"database/sql"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
	"log"
	"os"
	"strconv"
)

// Initiate logger
var logger *log.Logger

func main() {
	// Database connection
	dbSchemaParam := "loan_app"
	dbConnectionAddress := "user=postgres password=kopinikmat dbname=postgres sslmode=disable host=localhost port=5432"
	dbMaxOpenConnection := 500
	dbMaxIdleConnection := 100
	dBConnection, err := config.GetDbConnection(dbSchemaParam, dbConnectionAddress, dbMaxOpenConnection, dbMaxIdleConnection)
	if err != nil {
		fmt.Println("Error Connection Database")
		os.Exit(3)
	}

	// call sql migration
	dbMigrate(dBConnection)

	// call api controller
	router.ApiController()
}

func dbMigrate(dBConnection *sql.DB) {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}
	if dBConnection != nil {
		n, err := migrate.Exec(dBConnection, "postgres", migrations, migrate.Up)
		if err != nil {
			fmt.Println("Error migrations!", err.Error())
			logger.Println(err.Error())
			os.Exit(3)
		} else {
			//logger.Println("Applied " + strconv.Itoa(n) + " migrations!")
			fmt.Println("Applied " + strconv.Itoa(n) + " migrations!")
		}
	} else {
		logger.Println("empty database")
		fmt.Println("empty database")
		os.Exit(3)
	}
}
