package server

import (
	"fmt"
	"net/http"
	"time"

	"example.com/contact/internal/database"
)

type Server struct {
	db   database.Service
	port int
}

func NewServer() *http.Server {
	dbService := database.New()

	NewServer := &Server{
		db:   dbService,
		port: 5000,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
