package main

import (
	"fmt"                // package must be used after imported
	"learning-Go/helper" // need to explicitly tell Go it is not a built-in package
)

// package level variables
const conferenceTickets = 50

var conferenceName = "Go Conference" // synatic sugar: conferenceName := "Go Congerence"
var remainingTickets uint = 50       // unsigned integer, protect the variable from being negative
// var bookings [50]string   // need to assign the size for declaring arrays
var bookings = make([]map[string]string, 0) // slice: dynamic array in Go
// alternativelly, we can init slice like this:
// var bookings = []string{}, bookings := []string{}
// entry point of the application

func main() {

	fmt.Printf("conferenceTickets is %T, reremainingTickets is %T, confereconferenceName is %T\n",
		conferenceTickets, remainingTickets, conferenceName) // print out the data types of variables, %T is placeholder for data type

	helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets) // capitalize the first letter of the function to explicitly export the function from another pacakge

	// There is only one type of loop in Go: For loop
	for remainingTickets > 0 && len(bookings) < 50 {

		// ask for user input
		firstName, lastName, email, userTickets := helper.GetUserInput()

		// validate user input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets, bookings := helper.BookTickets(firstName, lastName, email, userTickets, remainingTickets, bookings, conferenceName)
			firstNames := helper.GetFirstNames(bookings)
			fmt.Printf("The first names of our bookings are %v\n", firstNames)

			// var noTicketsRemaining bool = remainingTickets == 0
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year")
				break // break the entire for loop
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}

	// switch syntax
	// city := "London"
	// switch city {
	// case "New York":
	// 	// execute code for booking New York conference tickets
	// case "Singapore", "London", "Berlin":
	// 	fmt.Println("London")
	// default:
	// 	fmt.Println("No valid city selected")
	// }
}
