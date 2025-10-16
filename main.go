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

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		//get user input
		firstName, lastName, emailAddress, userTickets := getUserInput()

		//validate the credentials that the user provided using the validateUserInput function
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//booking the ticket
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, emailAddress, conferenceName)

			//retrieve only the first name of the user bookings and print it to the user using the fucntion getFirstNames
			var firstNames = getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			//end the program if the remaining tickets is 0
			if remainingTickets == 0 {
				fmt.Println("Our Conference is fully booked. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered is incorrect")
			}
			if !isValidTicketNumber {
				fmt.Println("The ticket number you entered is incorrect")
			}

		}

	}

}

func greetUsers(confName string, confTickets int, confRemainTickets uint) {
	fmt.Printf("Welcome to our %v booking application!\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, confRemainTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, emailAddress string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(emailAddress, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {

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

	return firstName, lastName, emailAddress, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, emailAddress string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)

}
