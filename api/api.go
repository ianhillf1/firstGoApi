package api

import (
    "fmt"
    "net/http"
    "encoding/json"
	"github.com/gorilla/mux"
)

// BuildRouter returns a HTTP router that will handle requests appropriately
func BuildRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/customer/{cif}", GetCustomerDetails)
	return router
}

// GetCustomerDetails handles a HTTP request to fetch a customer matching the CIF in the request variables
func GetCustomerDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cif := vars["cif"]

	customer, err := CustomerRepo.GetCustomerDetails(cif)
	if(err != nil) {
		fmt.Fprint(w, err)
	} else {
		json.NewEncoder(w).Encode(customer)
	}
}