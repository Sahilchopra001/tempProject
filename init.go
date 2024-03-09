package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB initializes the database connection
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/your_database_name")
	if err != nil {
		return nil, err
	}
	return db, nil
}
