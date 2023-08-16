package main

import (
	"fmt"
	"strings"
)

func main() {

	confernceName := "Go Booking-Confernce"
	const confernceTicket = 50
	var remainingTicket uint = 50
	booking := []string{}

	fmt.Printf("welcome to %v\n", confernceName)
	fmt.Printf("We Have total of %v tickets and %v are still available.\n", confernceTicket, remainingTicket)
	fmt.Printf("Control get your ticket to attend\n")

	for {
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

		remainingTicket = remainingTicket - userTicket
		booking = append(booking, firstName+" "+lastName)

		fmt.Printf("Remaining ticket is equal %v\n", remainingTicket)
		fmt.Printf("Username is %v %v and email: %v Have been Booked %v tickets Thanks you!\n", firstName, lastName, email, userTicket)

		getFirstName := []string{}

		for _, book := range booking {
			var name = strings.Fields(book)
			getFirstName = append(getFirstName, name[0])
		}

		fmt.Printf("This are all the our booking : %v\n", booking)
		fmt.Printf("This are all the first name booking are : %v\n", getFirstName)
	}

}
