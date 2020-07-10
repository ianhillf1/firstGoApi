package repository

import (
	"time"
	"fmt"
)

// Customer is the structure defining the details of a customer
type Customer struct {
	Title string
	FirstName string
	Surname string
	DateOfBirth time.Time
	MobileNumber string
}

// JSONCustomer is used to deserialize the response from the API, for better
// control of date deserialization
type JSONCustomer struct {
	Title string
	FirstName string
	Surname string
	DateOfBirth string
	MobileNumber string
}

// CustomerRepo is the type which contains methods to get customer details
type CustomerRepo struct {
	BaseURL string
}

// GetCustomerDetails fetches a customer from the API
func (repo *CustomerRepo) GetCustomerDetails(cif string) (*Customer, error) {
	url:=fmt.Sprintf("%s/GetCustomerDetails?CustomerCif=%s", repo.BaseURL, cif)
	var customer JSONCustomer
	err := Client.GetJSONObject(url, nil, &customer)
	if err != nil {
		return nil, err
	}

	return mapCustomer(&customer)
}

func mapCustomer(customer *JSONCustomer) (*Customer, error) {
	dob, err := time.Parse("2006-01-02", customer.DateOfBirth)
	if err != nil {
		return nil, err
	}

	outputCustomer := Customer {
		FirstName: customer.FirstName,
		Surname: customer.Surname,	
		Title: customer.Title,	
		DateOfBirth: dob,		
		MobileNumber: customer.MobileNumber,
	}

	return &outputCustomer, nil
}