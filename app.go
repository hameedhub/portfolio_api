package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App start
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Init start the application configuration
func (a *App) Init(user, password, dbname string) {}

// Start the application with the address as param
func (a *App) Start(addr string) {}
