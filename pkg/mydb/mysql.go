package mydb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func SetUp() {
	log.Println("start to initialize mysql.")
	var err error
	DB, err = sql.Open("mysql", "root:abcd1234#$@%@tcp(127.0.0.1)/demo")

	if err != nil {
		log.Fatal(err)
	}
	log.Println("mysql initialized successfully.")
}

func Close() {
	if DB != nil {
		_ = DB.Close()
	}
}
