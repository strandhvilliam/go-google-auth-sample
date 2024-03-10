package utils

import (
	"google-auth-go-v2/internal/infra"
)

func SeedDb() error {
	db := infra.GetDb()

	query := `CREATE TABLE IF NOT EXISTS users (
						       id TEXT PRIMARY KEY NOT NULL,
						       first_name TEXT NOT NULL,
						       last_name TEXT NOT NULL,
                   email TEXT NOT NULL,
                   picture TEXT NOT NULL)`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
