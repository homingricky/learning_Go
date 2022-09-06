package main

import (
	"fmt" // package must be used after imported
)

// entry point of the application
func main() {
	var conferenceName = "Go Conference" // variables must be used after declared
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T, reremainingTickets is %T, confereconferenceName is %T",
		conferenceTickets, remainingTickets, conferenceName) // print out the data types of variables, %T is placeholder for data type

	fmt.Printf("Welcome to our %v booking application\n", conferenceName) // %v: placeholder for variable
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend") // println automatically add a newline for us at the end

	var userName string // it is required to explicitly define a data type if the value is not assigned at the beginning
	var userTickets int

	// ask the user for their name
	userName = "Ricky"
	userTickets = 2
	fmt.Printf("User %v booked %v tickers\n", userName, userTickets)
}
