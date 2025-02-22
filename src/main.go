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
	var email string
	var bookedTickets uint

	for {
		// Ask User info
		fmt.Println("\nUser form:")
		fmt.Print("Enter first name: ")
		fmt.Scan(&userFistName)
		fmt.Print("Enter last name: ")
		fmt.Scan(&userLastName)

		fmt.Print("Enter user email: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of desired tickets: ")
		fmt.Scan(&bookedTickets)

		remainingTickets -= bookedTickets
		bookings = append(bookings, userFistName+" "+userLastName)

		fmt.Printf("Thank you %v for buying %v tickets. You will receive your tickets at email %v\n",
			userFistName, email, bookedTickets)

		fmt.Printf("Remaining Tickets: %v\n", remainingTickets)
		fmt.Printf("The whole bookings array: %v\n", bookings)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("People that already booked tickets: %v\n", firstNames)

	}

}
