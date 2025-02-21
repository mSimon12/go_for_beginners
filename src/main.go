package main

import (
	"fmt"
)

func main() {
	movieTheaterName := "SuperCine"
	const theaterTotalTickets uint = 100
	var remainingTickets uint = theaterTotalTickets
	bookings := []string{}

	fmt.Println("Welcome to", movieTheaterName, "tickets Store!")
	fmt.Print("Book here your ticket to the most amazing movie experience.\n")
	fmt.Printf("Available tickets: %v\n", remainingTickets)

	var userName string
	var email string
	var bookedTickets uint

	// Ask User info
	fmt.Println("\nUser form:")
	fmt.Print("Enter user name: ")
	fmt.Scan(&userName)

	fmt.Print("Enter user email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of desired tickets: ")
	fmt.Scan(&bookedTickets)

	remainingTickets -= bookedTickets
	bookings = append(bookings, userName)

	fmt.Printf("Thank you %v for buying %v tickets. You will receive your tickets at email %v\n",
		userName, email, bookedTickets)

	fmt.Printf("Remaining Tickets: %v\n", remainingTickets)
	fmt.Printf("The whole bookings array: %v\n", bookings)
}
