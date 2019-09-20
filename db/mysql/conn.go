package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(47.94.207.144:3307)/file_store")
	if err != nil {
		os.Exit(1)
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		os.Exit(1)
	}
}

func GetDB() *sql.DB {
	return db
}
