package config

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
)

type DBInfo struct {
	instance      *sql.DB
	driver        string
	connectionStr string
	setParams     []string
}

func GetDbConnection(defaultSchema string, connectionString string, maxOpenConnection int, maxIdleConnection int) (*sql.DB, error) {
	// Attempt to establish a connection to the database
	db := connectToDatabase(connectionString)
	if db == nil {
		return nil, errors.New("failed to connect to the database")
	}

	// Check if the specified schema exists
	if !schemaExists(db, defaultSchema) {
		// If the schema doesn't exist, create it
		if err := createSchema(db, defaultSchema); err != nil {
			db.Close()
			return nil, err
		}
	}

	// Set connection pool configurations
	db.SetMaxOpenConns(maxOpenConnection)
	db.SetMaxIdleConns(maxIdleConnection)

	return db, nil
}

func connectToDatabase(connectionString string) *sql.DB {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		return nil
	}
	return db
}

func schemaExists(db *sql.DB, schemaName string) bool {
	query := "SELECT schema_name FROM information_schema.schemata WHERE schema_name = $1"
	var existingSchemaName string
	err := db.QueryRow(query, schemaName).Scan(&existingSchemaName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("Failed to check if schema exists: %v\n", err)
		return false
	}
	return existingSchemaName == schemaName
}

func createSchema(db *sql.DB, schemaName string) error {
	query := fmt.Sprintf("CREATE SCHEMA %s", schemaName)
	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("Failed to create schema: %v\n", err)
		return err
	}
	fmt.Printf("Schema %s created successfully\n", schemaName)
	return nil
}
