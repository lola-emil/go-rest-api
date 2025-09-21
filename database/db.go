package database

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMySQLStorage(cfg mysql.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, err
	}

	return db, nil
}
