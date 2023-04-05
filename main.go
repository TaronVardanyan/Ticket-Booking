package main
import "fmt"

func main() {
	confereceName := "Go Conference"
	const confereceTikets = 50
	remainingTickets := 50

	fmt.Printf("remainingTickets type is %T\n", remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", confereceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confereceTikets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Please enter how many tikets you wanna book: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n",firstName, lastName, userTickets, email)
}