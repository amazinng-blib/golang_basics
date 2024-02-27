package main

import (
	"fmt"
)

func main(){
	 name  := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTicket uint= 50

	fmt.Printf("conferenceTicket is  %T, remainingTicket is %T.\n", conferenceTickets, remainingTicket)

	fmt.Printf("Welcome to %v booking application\n", name)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTicket)
	fmt.Println("Get your tickets here to attend")
    
	var userName string
	var userTicket uint
	// ask user for thier name
    userName = "Tom"
	fmt.Printf("User %v booked %v tickets.\n", userName, userTicket)


 

}