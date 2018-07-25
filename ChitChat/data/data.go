package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "kento/limited7@/chitchat")
	if err != nil {
		log.Fatal(err)
	}

	return
}
