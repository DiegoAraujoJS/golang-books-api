package config

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Connect() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "samsung",
		Net:    "tcp",
		Addr:   "mysql",
		DBName: "books",
	}
	d, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err.Error)
	}
	pingErr := d.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
	fmt.Println("Connected to database" + "books")
	db = d
}

func GetDB() *sql.DB {
	return db
}
