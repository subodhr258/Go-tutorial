package helper

import "strings"

//capitalize the name of the function to export it.
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool,bool,bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

	//a Go function can return multiple values.
	return isValidName, isValidEmail, isValidTicketNumber
}

//scope of variables and functions: Local, Package Level, Global
