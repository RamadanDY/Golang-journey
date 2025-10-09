package main

import "fmt"

// OPERATOR
func main() {
	var remainingTickets uint = 50
	fmt.Println(remainingTickets)
	fmt.Println("this is the number of tickets that we have ", remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	// lets store the booking details in an Array

	var bookings [50]string
	// bookings[0] = firstName +"Nane"+ lastName

	fmt.Println("Enter your name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your lastName")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("enter the number of yh tickets")
	fmt.Scan(&userTicket)
	fmt.Printf("thank u %v %v for booking %v tickets.we will send u a confirmation via this email:%v\n", firstName, lastName, userTicket, email)
	bookings[0] = firstName + "  " + lastName
	fmt.Printf("lets print our the whole array%v\n", bookings)
	fmt.Printf("lets print the index %v\n", bookings[0])
	// fmt.Printf("lets print the pointer%v\n", &bookings)
	fmt.Printf("lets print the type %T\n", bookings)
	// this primts the size or len of the array
	fmt.Printf("lets print the type %v\n", len(bookings))

	remainingTickets = remainingTickets - userTicket

	fmt.Printf("we have %v remaining tickets", remainingTickets)
}
