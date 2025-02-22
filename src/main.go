package main

import (
	"fmt"
	"strings"
)

func main() {
	movieTheaterName := "SuperCine"
	const theaterTotalTickets uint = 100
	var remainingTickets uint = theaterTotalTickets
	bookings := []string{}

	fmt.Println("Welcome to", movieTheaterName, "tickets Store!")
	fmt.Print("Book here your ticket to the most amazing movie experience.\n")
	fmt.Printf("Available tickets: %v\n", remainingTickets)

	var userFistName string
	var userLastName string
	var userEmail string
	var desiredTickets uint

	for remainingTickets > 0 {
		// Ask User info
		fmt.Println("\nUser form:")
		fmt.Print("Enter first name: ")
		fmt.Scan(&userFistName)
		fmt.Print("Enter last name: ")
		fmt.Scan(&userLastName)

		fmt.Print("Enter user email: ")
		fmt.Scan(&userEmail)

		fmt.Print("Enter number of desired tickets: ")
		fmt.Scan(&desiredTickets)

		// Check inputted names
		isNameValid := len(userFistName) >= 2 && len(userLastName) >= 2

		// Check email
		isEmailValid := strings.Contains(userEmail, "@")

		// Check desired tickets
		isAmountValid := desiredTickets > 0 && desiredTickets <= remainingTickets

		if isNameValid && isEmailValid && isAmountValid {
			remainingTickets -= desiredTickets
			bookings = append(bookings, userFistName+" "+userLastName)

			fmt.Printf("Thank you %v for buying %v tickets. You will receive your tickets at email %v\n",
				userFistName, userEmail, desiredTickets)

			fmt.Printf("Remaining Tickets: %v\n", remainingTickets)
			fmt.Printf("The whole bookings array: %v\n", bookings)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
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
