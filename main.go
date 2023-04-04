package main
import "fmt"

func main() {
	var confereceName = "Go Conference"
	const confereceTikets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", confereceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confereceTikets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}