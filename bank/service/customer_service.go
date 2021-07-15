package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
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
		return nil, errs.NewUnexpectedError()
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
			return nil, errs.NewNotFoundError("customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	custResp := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &custResp, nil
}
