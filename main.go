package main

import (
	"fmt" // package must be used after imported
	"strings"
)

// package level variables
const conferenceTickets = 50

var conferenceName = "Go Conference" // synatic sugar: conferenceName := "Go Congerence"
var remainingTickets uint = 50       // unsigned integer, protect the variable from being negative
// var bookings [50]string   // need to assign the size for declaring arrays
var bookings []string // slice: dynamic array in Go
// alternativelly, we can init slice like this:
// var bookings = []string{}, bookings := []string{}
// entry point of the application

func main() {

	fmt.Printf("conferenceTickets is %T, reremainingTickets is %T, confereconferenceName is %T\n",
		conferenceTickets, remainingTickets, conferenceName) // print out the data types of variables, %T is placeholder for data type

	greetUsers() // function call

	// There is only one type of loop in Go: For loop
	for remainingTickets > 0 && len(bookings) < 50 {

		// ask for user input
		firstName, lastName, email, userTickets := getUserInput()

		// validate user input
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(firstName, lastName, email, userTickets)
			firstNames := getFirstNames()
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

// the func does not need parameters because package level variables are defined
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName) // %v: placeholder for variable
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend") // println automatically add a newline for us at the end
}

func getFirstNames() []string {
	firstNames := []string{}           // slice with empty initialization
	for _, booking := range bookings { // _: blank identifier: use to ignore a variable that are not used in Go
		var names = strings.Fields(booking) // splits the string with white space as separator and returns a slice with split elements
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= int(remainingTickets)
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, int) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, email string, userTickets int) {
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, lastName+" "+firstName)

	fmt.Printf("Thank you %v %v for booking %v tickers. You will receive a confirmation email at %v\n", lastName, firstName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)
}
