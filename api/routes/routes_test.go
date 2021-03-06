package routes

import (
	"time"
	"testing"
	"../../customer"
	"net/http"
	"net/http/httptest"
    "encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestGetCustomerDetailsCallsRepoCorrectly(t *testing.T) {
	mock := mockCustomerRepo{}
	customer.Repo = &mock

	request,_ := http.NewRequest("GET", "/customer/myTestCif123", nil)
	response := httptest.NewRecorder()
	myRouter := BuildRouter()

	myRouter.ServeHTTP(response, request)

	assert.EqualValues(t, 1, mock.requestCount, "Number of requests")
	assert.EqualValues(t, "myTestCif123", mock.lastRequestedCif, "Requested CIF")
}

func TestGetCustomerDetailsReturnsCustomerCorrectly(t *testing.T) {
	mock := mockCustomerRepo{}
	mock.customerToReturn = &customer.Customer {
		FirstName: "Ian",
		Surname: "Test",
		DateOfBirth: time.Date(1978, time.August, 17, 0, 0, 0, 0, time.UTC),
		Title: "Mr",
		MobileNumber: "07777123456",
	}
	customer.Repo = &mock

	request,_ := http.NewRequest("GET", "/customer/myTestCif123", nil)
	response := httptest.NewRecorder()
	myRouter := BuildRouter()

	myRouter.ServeHTTP(response, request)

	var result map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err, "Error was thrown by parsing JSON")
	assert.NotNil(t, result, "Parsed result as nil")
	assert.EqualValues(t, "Ian", result["FirstName"].(string), "FirstName")
	assert.EqualValues(t, "Test", result["Surname"].(string), "Surname")
	assert.EqualValues(t, "1978-08-17T00:00:00Z", result["DateOfBirth"].(string), "DateOfBirth")
	assert.EqualValues(t, "Mr", result["Title"].(string), "Title")
	assert.EqualValues(t, "07777123456", result["MobileNumber"].(string), "MobileNumber")
}

type mockCustomerRepo struct {
	requestCount int
	lastRequestedCif string
	customerToReturn *customer.Customer
	errorToReturn error
}

func (mock *mockCustomerRepo)GetCustomerDetails(cif string) (*customer.Customer, error) {
	mock.requestCount++
	mock.lastRequestedCif = cif
	return mock.customerToReturn, mock.errorToReturn
}
