# firstGoApi
# My first attempt at a REST API (essentially a passthrough) in Go.
# Reads customer details from the Test Data API on the tm dev environment and outputs them in JSON

Packages are:
- api  : contains the logic for responding to HTTP requests. Tested by tests/apitests
  - apiRunner  : contains the main entry point for running the api (on port 10000)
- consoletest  : contains a simple console tester to manually test the customer repository
- repository  : contains the repository layer with methods to fetch customer data from the dev Test Data API
- restclient  : abstraction layer to wrap the mechanics of a REST API call
- tests  : parent folder for unit tests
  - apitests  : unit tests for api package
  - repositorytests  : unit tests for repository package
  - restclienttests  : unit tests for restclient package
  
To run the API:
  go run .\api\apiRunner\apiRunner.go

To run the console test:
  go run .\consoletest\consoletest.go
  
To run the unit tests:
  go test .\tests\...
  
