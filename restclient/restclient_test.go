package restclient

import (
	"bytes"
	"io/ioutil"
	"testing"
	"net/http"
	
	"github.com/stretchr/testify/assert"
)

func TestGetJSONObjectMakesRequest(t *testing.T){
	restClient := RestClient { }
	mock := mockHTTPClient{ responseToReturn: BuildMockResponse("", 200) }
	Client = &mock

	var result map[string]interface{}
	var headers http.Header = http.Header{}
	restClient.GetJSONObject("http://myTestUrl/something", &headers, &result)

	assert.EqualValues(t, 1, mock.requestCount, "Number of requests")
	assert.NotNil(t, mock.lastRequest, "Request object")
	assert.EqualValues(t, "http://myTestUrl/something", mock.lastRequest.URL.String(), "URL")
	assert.Equal(t, mock.lastRequest.Header, headers, "Headers")
}

func TestGetJSONObjectParsesValidResponse(t *testing.T){
	restClient := RestClient { }
	mock := mockHTTPClient{ responseToReturn: BuildMockResponse(
		`{ "MyField":"MyValue", "MySubObject":{ "yarp":123, "narp":false}}`, 200) }
	Client = &mock

	var result map[string]interface{}
	var headers http.Header = http.Header{}
	err := restClient.GetJSONObject("http://myTestUrl/something", &headers, &result)

	assert.Nil(t, err, "Error")
	assert.NotNil(t, result, "Result")
	assert.Equal(t, "MyValue", result["MyField"], "MyField")
	assert.NotNil(t, result["MySubObject"], "MySubObject")
	assert.IsType(t, map[string]interface{}{}, result["MySubObject"], "MySubObject type")
	subObject := result["MySubObject"].(map[string]interface{})
	assert.EqualValues(t, 123, subObject["yarp"], "yarp")
	assert.EqualValues(t, false, subObject["narp"], "narp")
}

func BuildMockResponse(bodyText string, statusCode int) *http.Response {
	return &http.Response {
		StatusCode: statusCode,
		Body: ioutil.NopCloser(bytes.NewReader([]byte(bodyText))),
	}
}

type mockHTTPClient struct {
	requestCount int
	lastRequest *http.Request
	responseToReturn *http.Response
	errorToReturn error
}

func (mock *mockHTTPClient)Do(request *http.Request) (*http.Response, error) {
	mock.requestCount++
	mock.lastRequest = request
	return mock.responseToReturn, mock.errorToReturn
}