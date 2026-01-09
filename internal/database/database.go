package database

import (
	"fmt"
	"log"
	"time"

	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Service interface {
	GetInstance() *sqlx.DB

	Close() error
}

type service struct {
	db *sqlx.DB
}

var (
	dbname   = "go_contact"
	password = "789632145"
	username = "staleexam"
	host     = "localhost"
	port     = "3306"

	dbInstance *service
)

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname))

	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(50) // max total connections
	db.SetMaxIdleConns(25) // idle connections
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	dbInstance = &service{
		db: db,
	}

	return dbInstance
}

func (s *service) GetInstance() *sqlx.DB {
	return s.db
}

func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dbname)
	return s.db.Close()
}
