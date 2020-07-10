package customerrepotests

import (
	"testing"
	"../../repository"
	"net/http"
	
	"github.com/stretchr/testify/assert"
)

func TestGetCustomerDetailsCallsAPICorrectly(t *testing.T){
	customerRepo := repository.CustomerRepo{ BaseURL: "http://test.com/myApi"}
	mock := mockRestClient{ }
	repository.Client = &mock

	customerRepo.GetCustomerDetails("400600Test")

	assert.EqualValues(t, 1, mock.requestCount, "Number of requests")
	assert.EqualValues(t, "http://test.com/myApi/GetCustomerDetails?CustomerCif=400600Test", mock.lastRequestURL, "URL")
	var customer repository.JSONCustomer;
	assert.IsType(t, &customer, *mock.itemPassed, "Type of item passed")
}

type mockRestClient struct {
	requestCount int
	lastRequestURL string
	itemPassed *interface{}
	itemPopulator func(item interface{})
	errorToReturn error
}

func (mock *mockRestClient)GetJSONObject(url string, headers http.Header, item interface{}) (error){
	mock.requestCount++
	mock.lastRequestURL = url
	mock.itemPassed = &item
	if(mock.itemPopulator != nil) {
		mock.itemPopulator(item)
	}
	return mock.errorToReturn
}

