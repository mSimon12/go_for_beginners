package main

import "fmt"

func greetUsers() {
	fmt.Printf("Welcome to %v tickets Store!\n", movieTheaterName)
	fmt.Print("Book here your ticket to the most amazing movie experience.\n")
	fmt.Printf("Available tickets: %v\n", remainingTickets)
}
