package customer

import (
	"net/http"
	"../restclient"
)

// Repository is an interface for repositories which can fetch customer details
type Repository interface {
	GetCustomerDetails(cif string) (*Customer, error)
}

// RestClient is an interface for implementations of REST clients (can be mocked by setting Client)
type RestClient interface {
	GetJSONObject(url string, headers *http.Header, item interface{}) (error)
}

var (
	// Repo is the implementation of CustomerRepository to be used
	Repo Repository
	// Client is the implementation of RestClient to be used (can be mocked)
	Client RestClient
)

func init(){
	Repo = &RepoImpl { 
		BaseURL: "https://thinkmoney-dev.outsystemsenterprise.com/TMAutomatedTests_Api/rest/TestDataApi",
	}
	Client = &restclient.RestClient {}
}
