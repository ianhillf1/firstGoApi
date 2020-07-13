package routes

import (
	"github.com/gorilla/mux"
	"../../customer"
)

// BuildRouter returns a HTTP router that will handle requests appropriately
func BuildRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/customer/{cif}", customer.GetCustomerDetails)
	return router
}
