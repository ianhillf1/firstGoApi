package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

// HTTPClient is an interface which can be mocked
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	// Client can be set to mock
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

// RestClient is the type containing the methods to get and post Rest 
type RestClient struct {
}

// Get fetches a resource from the given URL
func (*RestClient) Get(url string, headers http.Header) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header = headers
	return Client.Do(request);
}

// GetJSONObject fetches a JSON-encoded object from the given URL
func (client *RestClient) GetJSONObject(url string, headers http.Header, item interface{}) (error) {
	response, err := client.Get(url, headers)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &item)
}

// Post posts a resource to the given URL
func (RestClient) Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers
	return Client.Do(request);
}