package helper

import (
	"fmt"
	"strings"
)


func ValidateInput() (bool, bool, string, string, string, int){
	var firstName string
	var userTickets int
	var lastName string
	var email string
	// ask user for thier name

	// todo: referencing user name using pointer
	
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter Number of tickets")
	fmt.Scan(&userTickets)


	isValidName :=len(firstName) > 2 && len(lastName) > 2
	isValidEmail :=strings.Contains(email, "@")

	return isValidEmail, isValidName, firstName, lastName, email, userTickets
}