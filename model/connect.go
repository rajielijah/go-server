package model

import (
	"database/sql"
	"fmt"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@(tcp:localhost:3306)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to the Database")
	return db
}
