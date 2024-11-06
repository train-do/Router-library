package database

import (
	"database/sql"
)

func InitDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=todo sslmode=disable password=superUser host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, err
}
