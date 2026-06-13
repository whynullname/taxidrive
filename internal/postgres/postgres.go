package postgres

import (
	"database/sql"
)

func NewConnection(address string) (*sql.DB, error) {
	db, err := sql.Open("pgx", address)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
