package main

import (
	"fmt"
	"strings"
)

func main() {
	var remainingTickets uint = 50
	fmt.Println(remainingTickets)
	fmt.Println("This is the number of tickets that we have:", remainingTickets)

	var bookings []string // define slice outside loop

	for {
		var firstName string
		var lastName string
		var email string
		var userTicket uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter the number of tickets:")
		fmt.Scan(&userTicket)

		fmt.Printf("Thank you %v %v for booking %v tickets. We will send a confirmation to: %v\n",
			firstName, lastName, userTicket, email)

		bookings = append(bookings, firstName+" "+lastName)

		firstNames := []string{}
		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These are all the bookings in the application:\n%v\n", firstNames)

		remainingTickets -= userTicket
		fmt.Printf("We have %v remaining tickets\n", remainingTickets)

		if remainingTickets == 0 {
			fmt.Println("All tickets are sold out. Thank you!")
			break
		}
	}
}
