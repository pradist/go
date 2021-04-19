package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	createTb := `CREATE TABLE IF NOT EXISTS pairs (
		DEVICE_ID	INTEGER	NOT NULL, 
		USER_ID		INTEGER NOT NULL
	);`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table pairs", err)
	}
	fmt.Printf("Create table success.\n")
}
