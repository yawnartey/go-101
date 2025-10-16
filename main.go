package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	emailAddress    string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {
		//get user input
		firstName, lastName, emailAddress, userTickets := getUserInput()

		//validate the credentials that the user provided using the validateUserInput function
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//booking the ticket
			bookTicket(userTickets, firstName, lastName, emailAddress)
			go sendTicket(userTickets, firstName, lastName, emailAddress)

			//retrieve only the first name of the user bookings and print it to the user using the fucntion getFirstNames
			var firstNames = getFirstNames()
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

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		emailAddress:    emailAddress,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending tickets: \n %v \n to email address %v \n", ticket, emailAddress)
	fmt.Println("##########")
}
