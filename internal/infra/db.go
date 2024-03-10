package infra

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sqlx.DB
}

var db = &Database{}

func (db *Database) Connect() (err error) {
	dbName := fmt.Sprintf("%s.db", os.Getenv("DB_NAME"))
	db.DB, err = sqlx.Open("sqlite3", dbName)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		defer db.Close()
		return fmt.Errorf("could not ping database: %w", err)
	}

	return nil
}

func GetDb() *Database {
	return db
}

func SeedDb() error {
	dbName := fmt.Sprintf("%s.db", os.Getenv("DB_NAME"))
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		file, err := os.Create(dbName)
		if err != nil {
			fmt.Println("Error creating database file:", err)
		}
		file.Close()
	}

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id TEXT PRIMARY KEY, email TEXT, first_name TEXT, last_name TEXT, picture TEXT)")
	if err != nil {
		return err
	}
	return nil
}
