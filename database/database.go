package database

import (
	"log"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func DBConnect(connStr string) {
	var err error
	DB, err = sqlx.Open("sqlite3", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
