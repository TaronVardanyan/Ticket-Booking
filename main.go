package main

import (
	"fmt"
	"strings"
)

const confereceTikets = 50
var confereceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := userInput()
		isValidName, isValidEmail, isValidTicketAmount := validateUser(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketAmount {
			bookTickets(userTickets, firstName, lastName, email)

			fmt.Printf("Firstnames of bookings are %v\n", getFirstNames())
			noTicketsCondition := remainingTickets == 0

			if noTicketsCondition {
				fmt.Println("All tickets was reserved")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address is not valid")
			}
			if !isValidTicketAmount {
				fmt.Println("Number of tickets that you entered is not valid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", confereceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confereceTikets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func userInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Please enter how many tikets you wanna book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n",firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confereceName)
}