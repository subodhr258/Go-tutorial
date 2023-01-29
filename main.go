package main

import (
	"booking-app/helper"
	"fmt"
)

//create global variables only if most of the functions need them.
//package level variables: variables defined at the top outside all functions.
var conferenceName = "Go Conference" 
// := is a shortcut for var. Can't be used for constants.
const conferenceTickets uint = 50
var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 0) //specify initial size of the slice
//or bookings := []string{}
//package level variables cannot be created using := syntax.

//create a custom type
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main() {
	//a slice is an abstraction of an array.
	//slices are more efficient than arrays

	//specifying a type explicitly is optional. 
	//It is only required when we want to create a type different from what Go would detect.
	//setting a uint type instead of int, will make sure we don't get a negative ticket count.

	greetUsers()

	//simplest infinite loop:
	for remainingTickets > 0 && len(bookings) < 50 {
		
		firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here.")
}

func getFirstNames() []string { //input and output parameters
	firstNames := []string{}
	for _, booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
	}
	// in go, _ are used to identify unuserd variables
	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// in go, all values have datatypes. reduces likelihood of errors.
	// go detects the errors even before we run the application

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //pass the memory address of that variable.
	//pointers are special variables in Go
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your first email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	// create a map for user
	var userData = UserData {
		firstName: firstName,
        lastName:  lastName,
        email:     email,
        numberOfTickets: userTickets,
	}

	bookings = append(bookings,userData)
	fmt.Printf("List of bookings is %v\n",bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}