package helper

import (
	"fmt"
	"strings"
)

// struct creation syntax
type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets int
}

// the func does not need parameters because package level variables are defined
func GreetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName) // %v: placeholder for variable
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend") // println automatically add a newline for us at the end
}

func GetFirstNames(bookings []UserData) []string {
	firstNames := []string{}           // slice with empty initialization
	for _, booking := range bookings { // _: blank identifier: use to ignore a variable that are not used in Go
		// var names = strings.Fields(booking) // splits the string with white space as separator and returns a slice with split elements
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= int(remainingTickets)
	return isValidName, isValidEmail, isValidTicketNumber
}

func GetUserInput() (string, string, string, int) {
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

func BookTickets(firstName string, lastName string, email string, userTickets int, remainingTickets uint, bookings []UserData, conferenceName string) (uint, []UserData) {
	remainingTickets = remainingTickets - uint(userTickets)

	// create a map for a user
	// var userData = make(map[string]string) // 1: data type of key; 2: data type of value
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10) // convert decimal uint into string

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickers. You will receive a confirmation email at %v\n", lastName, firstName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)
	return remainingTickets, bookings
}
