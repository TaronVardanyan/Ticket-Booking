package main

import (
	"fmt"
	"strings"
)

func main() {
	confereceName := "Go Conference"
	const confereceTikets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("remainingTickets type is %T\n", remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", confereceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confereceTikets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") && len(email) >= 2
		isValidTicketAmount := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketAmount {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			firstNames := []string{}

			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			noTicketsCondition := remainingTickets == 0

			if noTicketsCondition {
				fmt.Printf("All tickets was reserved\n")
				break
			}

			fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n",firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confereceName)
			fmt.Printf("Firstnames of bookings are %v\n", firstNames)
		} else {
			fmt.Printf("We only have %v tickets, so you can't book %v tickets \n", remainingTickets, userTickets)
		}
	}
}