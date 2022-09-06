package main

import (
	"fmt" //package must be used after imported
)

// entry point of the application
func main() {
	var conferenceName = "Go Conference" //variables must be used after declared
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets to attend")

}
