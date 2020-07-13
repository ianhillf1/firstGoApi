package customer

import (
	"time"
	"testing"
	"net/http"
	
	"github.com/stretchr/testify/assert"
)

func TestGetCustomerDetailsCallsAPICorrectly(t *testing.T){
	customerRepo := RepoImpl{ BaseURL: "http://test.com/myApi"}
	mock := mockRestClient{ }
	Client = &mock

	customerRepo.GetCustomerDetails("400600Test")

	assert.EqualValues(t, 1, mock.requestCount, "Number of requests")
	assert.EqualValues(t, "http://test.com/myApi/GetCustomerDetails?CustomerCif=400600Test", mock.lastRequestURL, "URL")
	var jsonCustomer jsonCustomer;
	assert.IsType(t, &jsonCustomer, *mock.itemPassed, "Type of item passed")
}

func TestGetCustomerDetailsReturnsCustomer(t *testing.T){
	customerRepo := RepoImpl{ BaseURL: "http://test.com/myApi"}
	mock := mockRestClient{ 
		itemPopulator: func(item interface{}) {
			jsonCustomer:=item.(*jsonCustomer)
			jsonCustomer.FirstName="IanTest"
			jsonCustomer.Surname="Testy"
			jsonCustomer.DateOfBirth="1976-09-25"
			jsonCustomer.Title="Wing Cmdr"
			jsonCustomer.MobileNumber="07777987654"
		},
	}
	Client = &mock

	customer,err := customerRepo.GetCustomerDetails("400600Test")

	assert.Nil(t, err, "Error returned")
	assert.NotNil(t, customer, "No customer returned")
	assert.EqualValues(t, "IanTest", customer.FirstName, "FirstName")
	assert.EqualValues(t, "Testy", customer.Surname, "Surname")
	assert.EqualValues(t, "Wing Cmdr", customer.Title, "Title")
	assert.EqualValues(t, "07777987654", customer.MobileNumber, "MobileNumber")
	assert.EqualValues(t, time.Date(1976, time.September, 25, 0, 0, 0, 0, time.UTC), customer.DateOfBirth, "DateOfBirth")
}

type mockRestClient struct {
	requestCount int
	lastRequestURL string
	itemPassed *interface{}
	itemPopulator func(item interface{})
	errorToReturn error
}

func (mock *mockRestClient)GetJSONObject(url string, headers *http.Header, item interface{}) (error){
	mock.requestCount++
	mock.lastRequestURL = url
	mock.itemPassed = &item
	if(mock.itemPopulator != nil) {
		mock.itemPopulator(item)
	}
	return mock.errorToReturn
}

