package main

import (
	"bank/repository"
	"bank/service"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:P@ssw0rd@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	custRepo := repository.NewCustomerRepository(db)
	custService := service.NewCustomerService(custRepo)
	customers, err := custService.GetCustomers()
	if err != nil {
		panic(err)
	}
	fmt.Println(customers)
}
