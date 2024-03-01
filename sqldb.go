package core

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewSqlDB creates a new connector to SQL Database
func NewSqlDB(host, port, user, password, dbName, driverName string) (*sqlx.DB, error) {
	// Generating a connection string
	if driverName == "" {
		driverName = "postgres" // Default driver
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	// Connecting to the database
	db, err := sqlx.Connect(driverName, dsn)
	if err != nil {

		return nil, fmt.Errorf("не удалось подключиться к %s: %v", host, err)
	}

	return db, nil
}
