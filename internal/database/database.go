package database

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	GetInstance() *sqlx.DB

	Close() error
}

type service struct {
	db *sqlx.DB
}

var (
	dbname   = os.Getenv("DB_NAME")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME") //"staleexam"
	host     = os.Getenv("DB_HOST")     // "localhost"
	port     = os.Getenv("DB_PORT")     // "3306"

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
