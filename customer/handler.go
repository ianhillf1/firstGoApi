package customer

import (
    "fmt"
    "net/http"
    "encoding/json"
	"github.com/gorilla/mux"
)

// GetCustomerDetails handles a HTTP request to fetch a customer matching the CIF in the request variables
func GetCustomerDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cif := vars["cif"]

	customer, err := Repo.GetCustomerDetails(cif)
	if(err != nil) {
		fmt.Fprint(w, err)
	} else {
		json.NewEncoder(w).Encode(customer)
	}
}