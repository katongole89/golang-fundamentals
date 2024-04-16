package main

import (
	"booking-app/helper"
	"fmt"
)

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
	// var testArray [50]string --- Just defining
	//empty array
	var bookings = [50]string{}
	// to add element
	bookings[0] = "call me"

	// better option is Slice - its an abstraction of an array
	// var testSlice = []string{}
	testSlice := []string{}
	testSlice = append(testSlice, "the string added")

	fmt.Println(testSlice)
	fmt.Println(bookings)

	// Loop that never ends
	// for {
	// }

	// Immitation of the while loop
	// for x>2{

	// }

	namesArray := []string{"rick", "durant", "lebron", "calton"}

	// Underscores are called blank identifiers - we expect a value eg index but we are not using it
	for _, name := range namesArray {
		fmt.Println(name)
		if name == "lebron" {
			break
		}

		// #use continue while in the for loop but u want to skip code

	}

	// IF statement
	if len(namesArray) > 0 {
		fmt.Println("WE ARE JUST WINNING")
	} else if len(namesArray) > 2 {
		fmt.Println("WE ALWAYS REACH HERE")
	} else {
		fmt.Println("WE FAILED")
	}

	//Switch statement
	switch conferenceName {
	case "kendrick":
		fmt.Println("WE FAILED")
	case "Drake":
		fmt.Println("WE FAILED")
	default:
		fmt.Println("J COLE")
	}

	greetUsers(conferenceName)
	daton := helper.ReturnArrayUsers(namesArray)
	fmt.Println(daton)

	var arrayUserData = make([]map[string]string, 0)
	fmt.Println(arrayUserData)

	var userData = make(map[string]string)
	userData["name"] = "katongole"
	userData["age"] = "24"
	userData["email"] = "katon@gmail.com"

	arrayUserData = append(arrayUserData, userData)
	fmt.Println(arrayUserData)

}

func greetUsers(name string) {
	fmt.Printf("Welcome %v", name)
}

// Returning function
// func returnArrayUsers(theData []string) []string {
// 	return theData
// }

//function that returns multiple values
// func retMultipleVal(name string)(bool, bool, []string){
// 	return val1, val2, val3
// }
// USING
