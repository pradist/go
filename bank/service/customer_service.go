package service

import (
	"bank/logs"
	"bank/repository"
	"database/sql"
	"errors"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{
		custRepo: custRepo,
	}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
	}
	custResps := []CustomerResponse{}
	for _, cust := range customers {
		custResp := CustomerResponse{
			CustomerID: cust.CustomerID,
			Name:       cust.Name,
			Status:     cust.Status,
		}
		custResps = append(custResps, custResp)
	}
	return custResps, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		logs.Error(err)
		return nil, err
	}
	custResp := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &custResp, nil
}
