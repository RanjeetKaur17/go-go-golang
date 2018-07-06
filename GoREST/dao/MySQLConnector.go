package dao

import (
	_ "github.com/go-sql-driver/mysql"
	. "database/sql"
	"time"
        "fmt"
)

var db *DB

/*
 * Initialize database connection
 */
func InitializeMySQL() {
	dBConnection, err := Open("mysql", "root:password@(localhost:3306)/student")
	if err != nil {
	    fmt.Println("Connection Failed!!")
	}
        err = dBConnection.Ping()
        if err != nil {
            fmt.Println("Ping Failed!!")
        }
	db = dBConnection
	dBConnection.SetMaxOpenConns(10)
	dBConnection.SetMaxIdleConns(5)
	dBConnection.SetConnMaxLifetime(time.Second * 10)
}

/*
 * Get database connection
 */
func GetMySQLConnection() *DB {
	return db
}

