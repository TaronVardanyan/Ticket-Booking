package main

import "strings"

func validateUser(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && len(email) >= 2
	isValidTicketAmount := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketAmount
}