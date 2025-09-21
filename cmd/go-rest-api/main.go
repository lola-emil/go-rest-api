package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/lola-emil/go-rest-api/database"
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

	_, err := database.NewMySQLStorage(cfg)

	if err != nil {
		log.Fatalf("Database Error: %s", err.Error())
	}

}
