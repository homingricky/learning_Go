package main

import (
	"fmt" // package must be used after imported
)

// entry point of the application
func main() {
	var conferenceName = "Go Conference" // synatic sugar: conferenceName := "Go Congerence"
	const conferenceTickets = 50
	var remainingTickets uint = 50 // unsigned integer, protect the variable from being negative

	fmt.Printf("conferenceTickets is %T, reremainingTickets is %T, confereconferenceName is %T",
		conferenceTickets, remainingTickets, conferenceName) // print out the data types of variables, %T is placeholder for data type

	fmt.Printf("Welcome to our %v booking application\n", conferenceName) // %v: placeholder for variable
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend") // println automatically add a newline for us at the end

	var firstName string // it is required to explicitly define a data type if the value is not assigned at the beginning
	var lastName string
	var email string
	var userTickets int

	// ask the user info
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // pass the address of the variable so Go could assign user inputted value back to the variable

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the no. of tickets you want to book")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickers. You will receive a confirmation email at %v\n", lastName, firstName, userTickets, email)

	remainingTickets = remainingTickets - uint(userTickets)
	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)

}
