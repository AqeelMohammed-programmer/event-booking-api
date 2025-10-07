package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const MAX_OPEN_CONNECTIONS = 10
const MAX_IDLE_CONNECTIONS = 5

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		log.Fatalln("ERROR: cannot initialize the database")
	}

	DB.SetMaxOpenConns(MAX_OPEN_CONNECTIONS)
	DB.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
}
