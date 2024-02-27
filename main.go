package main

import (
	"fmt"
	"strings"
)

func main(){
	 name  := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTicket uint= 50

	fmt.Printf("conferenceTicket is  %T, remainingTicket is %T.\n", conferenceTickets, remainingTicket)

	fmt.Printf("Welcome to %v booking application\n", name)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTicket)
	fmt.Println("Get your tickets here to attend")

	// todo: cfreating arrays

	// var bookings [50]string

	// todo: slices
	var bookings []string

	for {
		var firstName string
		var userTickets uint
		var lastName string
		var email string
		// ask user for thier name

		// todo: referencing user name using pointer
		
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("Enter Number of tickets")
		fmt.Scan(&userTickets)


		// todo: check if userTicket is greater than remaining ticket
	
		if userTickets > remainingTicket {
			fmt.Printf("Impossible, you can only book %v number of tickets\n", remainingTicket)
			
			// todo: continue specifies that the loop should go again if the first one fail
			continue
		}

		remainingTicket = remainingTicket -userTickets
		// bookings[0] = firstName + " " + lastName
		// todo: using slice
	

		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n",firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTicket, name)

	   firstNames :=[]string{}
	   for _, booking := range bookings{
		   var names = strings.Fields(booking)
		 
           firstNames = append(firstNames, names[0])
	   } 

		fmt.Printf("The first names of all our bookings are : %v\n", firstNames)
        
	
          
		if remainingTicket == 0{
			// todo: quit the app
			fmt.Println("Our conference is booked-out.Come back next year")
			break
		}
	}

}