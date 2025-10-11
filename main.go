package main

import (
	"fmt"
)

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint

	//ask the user for their booking details
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&emailAddress)

	fmt.Println("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)

	//calculate the ramining tickets after user booked their ticket
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("%v tickets remaing for %v", remainingTickets, conferenceName)
}
