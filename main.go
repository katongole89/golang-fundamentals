package main

import "fmt"

func main() {
	// var conferenceName = "Go conference"
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets int = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// print the types of variables
	fmt.Printf(" hey %T\n", remainingTickets)

	// var userName string
	// userName = "katon"
	// fmt.Println(userName)

	var userName string
	var userTickets int

	fmt.Printf("Enter your name\n")

	fmt.Scan(&userName)
	fmt.Printf("Enter the number of tickets\n")
	fmt.Scan(&userTickets)

	fmt.Printf("%v .u have bought %v tickets", userName, userTickets)

	// Arrays
	var testArray [50]string
	var bookings = [50]string{}
	fmt.Println(bookings)

}
