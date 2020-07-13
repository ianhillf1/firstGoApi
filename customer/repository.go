package customer

import (
	"net/http"
	"time"
	"fmt"
)

// jsonCustomer is used to deserialize the response from the API, for better
// control of date deserialization.
type jsonCustomer struct {
	Title string
	FirstName string
	Surname string
	DateOfBirth string
	MobileNumber string
}

// RepoImpl is the type which contains methods to get customer details
type RepoImpl struct {
	BaseURL string
}

// GetCustomerDetails fetches a customer from the API
func (repo *RepoImpl) GetCustomerDetails(cif string) (*Customer, error) {
	url:=fmt.Sprintf("%s/GetCustomerDetails?CustomerCif=%s", repo.BaseURL, cif)
	var jsonCustomer jsonCustomer
	err := Client.GetJSONObject(url, &http.Header{}, &jsonCustomer)
	if err != nil {
		return nil, err
	}

	return mapCustomer(&jsonCustomer)
}

func mapCustomer(jsonCustomer *jsonCustomer) (*Customer, error) {
	dob, err := time.Parse("2006-01-02", jsonCustomer.DateOfBirth)
	if err != nil {
		return nil, err
	}

	outputCustomer := Customer {
		FirstName: jsonCustomer.FirstName,
		Surname: jsonCustomer.Surname,	
		Title: jsonCustomer.Title,	
		DateOfBirth: dob,		
		MobileNumber: jsonCustomer.MobileNumber,
	}

	return &outputCustomer, nil
}