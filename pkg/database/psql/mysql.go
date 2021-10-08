package psql

import (
	"database/sql"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func Connect(URL string) (*sql.DB, error) {
	// db, err := sql.Open("mysql", Details)
	db, err := sql.Open("postgres", URL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
