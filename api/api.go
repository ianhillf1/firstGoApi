package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
	"github.com/gorilla/mux"
	//"../repository"
)

func main() {
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/customer/{cif}", getCustomerDetails)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func getCustomerDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cif := vars["cif"]

	customer, err := CustomerRepo.GetCustomerDetails(cif)
	if(err != nil) {
		fmt.Fprint(w, err)
	} else {
		json.NewEncoder(w).Encode(customer)
	}
}