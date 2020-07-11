package repository

import (
	"net/http"
	"../restclient"
)

// RestClient is an interface for implementations of REST clients (can be mocked by setting Client)
type RestClient interface {
	GetJSONObject(url string, headers *http.Header, item interface{}) (error)
}

var (
	// Client is the implementation of RestClient to be used (can be mocked)
	Client RestClient
)

func init() {
	Client = &restclient.RestClient {}
}

