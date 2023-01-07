package main

import (
	"database/sql"
	"fmt"
	// pg "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 8585
	user     = "postgres"
	password = "root"
	dbname   = "restdb"
)

func connectionpgSQL() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}

func main() {

}
