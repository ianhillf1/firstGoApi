# firstGoApi

My first attempt at a REST API (essentially a passthrough) in Go.
Reads customer details from the Test Data API on the tm dev environment and outputs them in JSON

Packages are:
- api  : contains the main entry point for running the api (on port 10000)
  - routes*  : contains the routing logic to handle individual api requests
- consoletest  : contains a simple console tester to manually test the customer repository
- customer*  : contains the handler and repository layer with methods to fetch customer data from the dev Test Data API
- restclient*  : abstraction layer to wrap the mechanics of a REST API call

* contains unit tests
  
To run the API:
  go run .\api
and then visit
  http://localhost:10000/customer/4006001200 (or other customer number)

To run the console test:
  go run .\consoletest
  
To run the unit tests:
  go test .\...
  
