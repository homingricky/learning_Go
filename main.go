package main

import (
	"fmt" // package must be used after imported
	"strings"
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

	// var bookings [50]string   // need to assign the size for declaring arrays
	var bookings []string // slice: dynamic array in Go
	// alternativelly, we can init slice like this:
	// var bookings = []string{}, bookings := []string{}

	// There is only one type of loop in Go: For loop
	for {
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

		if userTickets <= int(remainingTickets) {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, lastName+" "+firstName)

			fmt.Printf("Thank you %v %v for booking %v tickers. You will receive a confirmation email at %v\n", lastName, firstName, userTickets, email)
			fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings { // _: blank identifier: use to ignore a variable that are not used in Go
				var names = strings.Fields(booking) // splits the string with white space as separator and returns a slice with split elements
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of our bookings are %v", firstNames)

			// var noTicketsRemaining bool = remainingTickets == 0
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year")
				break // break the entire for loop
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, please book again", remainingTickets)
		}

	}
}
