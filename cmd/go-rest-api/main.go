package main

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lola-emil/go-rest-api/database"
	"github.com/lola-emil/go-rest-api/internals/api"
	"github.com/lola-emil/go-rest-api/internals/configs"
)

func main() {
	cfg := mysql.Config{
		User:                 configs.Env.DBUser,
		Passwd:               configs.Env.DBPassword,
		Addr:                 configs.Env.DBAddress,
		DBName:               configs.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := database.NewMySQLStorage(cfg)

	if err != nil {
		log.Fatalf("Database Error: %s", err.Error())
	}

	initStorage(db)

	server := api.NewApiServer(fmt.Sprintf(":%s", configs.Env.Port), db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sqlx.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
