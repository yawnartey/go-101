package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	for {

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

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			//store users booking records
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
			fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)

			//retrieve only the first name of the user bookings and print it out to the user
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings we have are: %v\n", firstNames)

			//end the program if the remaining tickets is 0
			if remainingTickets == 0 {
				fmt.Println("Our Conference is fully booked. Come back next year")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining. You can't book %v tickets. Enter a valid ticket number less or equal to %v\n", remainingTickets, userTickets, remainingTickets)
		}

	}

}
