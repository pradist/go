package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func main() {
	ds := fmt.Sprintf("user=\"%s\" password=\"%s\" connectString=\"%s:%d/%s\"",
		"ora_pf",
		"ora_pf",
		"cuatlrphubdb-scan",
		1525,
		"PHUBDB")
	fmt.Println(ds)
	db, err := sql.Open("godror", ds)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select COMP_ID from BILL_PAYMENTS")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		compID := ""
		rows.Scan(&compID)
		fmt.Println(compID)
	}

	// var v1 string
	// for rows.Next() {
	// 	rows.Scan(&v1)
	// }
	// fmt.Printf("Data: %s\n", v1)

	// tx, err := db.Begin()

	// query := "BEGIN " +
	// 	"INSERT INTO COMP_CODE_GENERATOR N " +
	// 	"(N.COMP_CODE, N.STATUS) " +
	// 	"VALUES (:val1, :val2); " +
	// 	"END;"

	// stmt, err := tx.Prepare(query)

	// for i := 0; i < 100000; i++ {

	// 	fmt.Println("No. ", i)
	// 	code := fmt.Sprintf("%05d", i)

	// 	var params = []interface{}{
	// 		code, "N",
	// 	}

	// 	if _, err := stmt.ExecContext(context.Background(), params...); err != nil {
	// 		tx.Rollback()
	// 		fmt.Println("Error save", err)
	// 		return
	// 	}
	// }

	// tx.Commit()

	//var thedate string
	//for rows.Next() {
	//
	//	rows.Scan(&thedate)
	//}
	//fmt.Printf("The date is: %s\n", thedate)
}
