package main

import (
	"fmt"
	"strings"
)

// OPERATOR
func main() {
	var remainingTickets uint = 50
	fmt.Println(remainingTickets)
	fmt.Println("this is the number of tickets that we have ", remainingTickets)
	for {

		var firstName string
		var lastName string
		var email string
		var userTicket uint

		// lets store the booking details in an Array
		// this is how to store values in an array
		// but note to put both the savings at the bottom
		// var bookings [50]string
		// bookings[0] = firstName +"Nane"+ lastName

		// this is how its done in slice
		var bookings []string
		// bookings = append(bookings, firstName+" "+lastName)

		fmt.Println("Enter your name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your lastName")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		fmt.Println("enter the number of yh tickets")
		fmt.Scan(&userTicket)
		fmt.Printf("thank u %v %v for booking %v tickets.we will send u a confirmation via this email:%v\n", firstName, lastName, userTicket, email)
		// bookings[0] = firstName + "  " + lastName
		bookings = append(bookings, firstName+" "+lastName)
		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("these are all the bookings in the application \n%v\n", firstNames)
		remainingTickets = remainingTickets - userTicket

		fmt.Printf("we have %v remaining tickets\n", remainingTickets)
	}
}
