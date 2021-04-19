package main

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

func main() {
	ds := fmt.Sprintf("user=\"%s\" password=\"%s\" connectString=\"%s:%d/%s\"", "ora_pf", "ora_pf", "172.30.74.93", 1525, "PHUDB")
	db, err := sql.Open("godror", ds)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}
