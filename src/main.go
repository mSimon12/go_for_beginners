package main

import (
	"fmt"
	"go_app/src/helper"
	"strconv"
)

const theaterTotalTickets uint = 100

var movieTheaterName string = "SuperCine"
var remainingTickets uint = theaterTotalTickets
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for remainingTickets > 0 {
		// Ask User info
		userFistName, userLastName, userEmail, desiredTickets := helper.GetUserInput()
		isNameValid, isEmailValid, isAmountValid := helper.ValidateUserInput(userFistName, userLastName, userEmail, desiredTickets, remainingTickets)
		if isNameValid && isEmailValid && isAmountValid {

			bookTicket(userFistName, userLastName, userEmail, desiredTickets)

			fmt.Printf("Remaining Tickets: %v\n", remainingTickets)
			fmt.Printf("The whole bookings array: %v\n", bookings)

			firstNames := extractBookingsFirstName()
			fmt.Printf("People that already booked tickets: %v\n", firstNames)

		} else {
			fmt.Println()
			if !isNameValid {
				fmt.Println("Invalid name!")
			}
			if !isEmailValid {
				fmt.Println("Invalid email!")
			}
			if !isAmountValid {
				fmt.Println("Invalid amount!")
			}
			fmt.Println("Invalid input data! Try again.")
		}

	}

	fmt.Println("\nThe desired movie session is sold out!")
}

func extractBookingsFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = booking["firstName"]
		firstNames = append(firstNames, names)
	}
	return firstNames
}

func bookTicket(firstName string, lastName string, email string, desiredTickets uint) {
	remainingTickets -= desiredTickets

	// create a map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(desiredTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v for buying %v tickets. You will receive your tickets at email %v\n",
		firstName, email, desiredTickets)
}
