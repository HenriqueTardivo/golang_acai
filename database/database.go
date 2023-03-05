package database

import "database/sql"

var DB *sql.DB

func Connect() error {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
