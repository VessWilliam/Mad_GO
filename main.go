package main

import (
	"fmt"
	"strconv"
)

var confernceName = "Golang Booking-App"

const confernceTicket int = 50

var remainingTicket uint = 50
var booking = make([]map[string]string, 0)

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTicket := getUserInput()

		isValidName, isValidEmail, isValidTicket := ValidateUserInput(firstName, lastName, email, userTicket)

		if isValidName && isValidEmail && isValidTicket {

			bookTicke(userTicket, firstName, lastName, email)

			getFirstName := getFirstNams()
			fmt.Printf("This are all the first name booking are : %v\n", getFirstName)

			if remainingTicket == 0 {
				fmt.Printf("Opss the the ticket have book out, please purchase next time")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("your last name and first name enter is too shot")
			}
			if !isValidEmail {
				fmt.Println("email enter need contain symbol @ sign")
			}
			if !isValidTicket {
				fmt.Println("number of tickets you enter is invalid")
			}
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to %v\n", confernceName)
	fmt.Printf("We Have total of %v tickets and %v are still available.\n", confernceTicket, remainingTicket)
	fmt.Printf("Control get your ticket to attend\n")
}

func getFirstNams() []string {
	getFirstName := []string{}
	for _, book := range booking {
		getFirstName = append(getFirstName, book["firstName"])
	}

	return getFirstName
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your Ticket purchase: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicke(userTicket uint, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTicket

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["ticket"] = strconv.FormatUint(uint64(userTicket), 10)

	booking = append(booking, userData)

	fmt.Printf("Remaining ticket is equal %v\n", remainingTicket)
	fmt.Printf("Username is %v %v and email: %v Have been Booked %v tickets Thanks you!\n", firstName, lastName, email, userTicket)
}
