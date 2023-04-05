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

	var userName string
	var userTickets int

	//some logic

	userName = "Tom"
	userTickets = 8
	fmt.Println(userName, userTickets)
}