package mysql

import (
	"database/sql"
	"os"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	// db, err := sql.Open("mysql", Details)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
