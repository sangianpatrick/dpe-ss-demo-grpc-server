package database

import (
	"database/sql"

	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetDatabase() *sql.DB {
	if db == nil {
		db = newDb()
	}

	return db
}

func newDb() *sql.DB {
	c := config.GetConfig()
	db, _ := sql.Open("mysql", c.Mariadb.DSN)

	return db
}
