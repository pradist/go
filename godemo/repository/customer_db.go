package repository

import "fmt"

type customerRepositoryDB struct {
	db DB
}

func NewCustomerRepository(db DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := `select customer_id, name, date_of_birth, city, zipcode, status 
		from customers`
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := `select customer_id, name, date_of_birth, city, zipcode, status 
		from customers 
		where customer_id = ?`
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r customerRepositoryDB) Insert(customer Customer) (*Customer, error) {
	query := `insert into customers (name, date_of_birth, city, zipcode, status)
		values (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, customer.Name, customer.DateOfBirth, customer.City, customer.ZipCode, customer.Status)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	fmt.Println(id)
	customer.CustomerID = int(id)
	return &customer, nil
}
