package main

import (
	"fmt"
	"sync"
	"time"
)

var confernceName = "Golang Booking-App"

const confernceTicket int = 50

var remainingTicket uint = 50
var booking = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	numOfTicket uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	for {
		firstName, lastName, email, userTicket := getUserInput()

		isValidName, isValidEmail, isValidTicket := ValidateUserInput(firstName, lastName, email, userTicket)

		if isValidName && isValidEmail && isValidTicket {

			bookTicke(userTicket, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTicket, firstName, lastName, email)

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
		getFirstName = append(getFirstName, book.firstName)
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

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		numOfTicket: userTicket,
	}

	booking = append(booking, userData)

	fmt.Printf("Print List of booking is %v\n", booking)

	fmt.Printf("Remaining ticket is equal %v\n", remainingTicket)
	fmt.Printf("Username is %v %v and email: %v Have been Booked %v tickets Thanks you!\n", firstName, lastName, email, userTicket)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for  %v %v", userTickets, firstName, lastName)
	fmt.Printf("**********************")
	fmt.Printf("Sending tiket:\n %v \nto email address %v\n", ticket, email)
	fmt.Printf("***********************")
	wg.Done()
}
