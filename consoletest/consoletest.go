package main

import (
	"fmt"
	"../repository"
)

const dateFormatOutput string = "02/01/2006"

func main() {
	repo := repository.CustomerRepo{ 
		BaseURL: "https://thinkmoney-dev.outsystemsenterprise.com/TMAutomatedTests_Api/rest/TestDataApi",
	}
	customer, err := repo.GetCustomerDetails("4006001200")
	if(err != nil){
		fmt.Println("Error fetching customer:", err)
	} else {
		fmt.Println(
			customer.Title,
			customer.FirstName, 
			customer.Surname, "was born on", 
			customer.DateOfBirth.Format(dateFormatOutput),
			"and can be contacted on",
			customer.MobileNumber)
	}
}