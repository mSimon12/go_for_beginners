package helper

import (
	"fmt"
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// Check inputted names
	isNameValid := len(firstName) >= 2 && len(lastName) >= 2

	// Check email
	isEmailValid := strings.Contains(email, "@")

	// Check desired tickets
	isAmountValid := userTickets > 0 && userTickets <= remainingTickets

	return isNameValid, isEmailValid, isAmountValid
}

func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var tickets uint

	// Ask User info
	fmt.Println("\nUser form:")
	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter user email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of desired tickets: ")
	fmt.Scan(&tickets)

	return firstName, lastName, email, tickets
}
