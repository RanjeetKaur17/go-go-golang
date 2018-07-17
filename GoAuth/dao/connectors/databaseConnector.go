package connectors

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

//Connect to database using provided credentials
func InitializeDataBase(host string, port string, user string, password string, dbname string){

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dBConnection, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Errorf("Database Connection Failed!!")
	}
	err = dBConnection.Ping()
	if err != nil {
		fmt.Errorf("Database Ping Failed!!")
	}
	db = dBConnection
	dBConnection.SetMaxOpenConns(100)
	dBConnection.SetMaxIdleConns(50)
	dBConnection.SetConnMaxLifetime(time.Second * 10)

	fmt.Println("Database Connection Successful")
}

func GetDatabaseConnection() *sql.DB {
	return db
}
