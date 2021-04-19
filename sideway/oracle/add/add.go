package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/godror/godror"
)

func worker(id int, jobs <-chan int, results chan<- int, tx *sql.Tx, stmt *sql.Stmt) {
	for j := range jobs {
		code := fmt.Sprintf("%05d", j)
		fmt.Println("worker", id, "started  job", code)

		var params = []interface{}{
			code, "N",
		}

		if _, err := stmt.ExecContext(context.Background(), params...); err != nil {
			tx.Rollback()
			fmt.Println("Error save", err)
			return
		}

		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", code)
		results <- j * 2
	}
}

func main() {

	ds := fmt.Sprintf("user=\"%s\" password=\"%s\" connectString=\"%s:%d/%s\"", "", "", "", 1521, "")
	fmt.Println(ds)
	db, err := sql.Open("godror", ds)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected")
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	query := "BEGIN " +
		"INSERT INTO COMP_CODE_GENERATOR2 N " +
		"(N.COMP_CODE, N.STATUS) " +
		"VALUES (:val1, :val2); " +
		"END;"

	stmt, err := tx.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}

	const numJobs = 10000
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 2000; w++ {
		go worker(w, jobs, results, tx, stmt)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

	tx.Commit()
}
