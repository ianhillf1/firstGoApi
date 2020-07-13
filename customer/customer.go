package customer

import (
	"time"
)

// Customer is the structure defining the details of a customer
type Customer struct {
	Title string
	FirstName string
	Surname string
	DateOfBirth time.Time
	MobileNumber string
}