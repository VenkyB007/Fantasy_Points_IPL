package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	password = "1234"
	port     = 5432
	user     = "postgres"
	dbname   = "player"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	fmt.Println("db connected successfully")
	return db
}
