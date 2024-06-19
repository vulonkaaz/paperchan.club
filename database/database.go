package database

import (
	"log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func DBConnect(connStr string) {
	var err error
	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
