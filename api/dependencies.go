package api

import (
	"../repository"
)

// CustomerRepository is an interface for repositories which can fetch customer details
type CustomerRepository interface {
	GetCustomerDetails(cif string) (*repository.Customer, error)
}

var (
	// CustomerRepo is the implementation of CustomerRepository to be used
	CustomerRepo CustomerRepository
)

func init(){
	CustomerRepo = &repository.CustomerRepo { 
		BaseURL: "https://thinkmoney-dev.outsystemsenterprise.com/TMAutomatedTests_Api/rest/TestDataApi",
	}
}