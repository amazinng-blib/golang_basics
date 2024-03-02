package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

// todo: struct

	type UserDataStructure struct {
		firstName string
		lastName string
		email string
		NumOfTickets uint
		isOptedInForNewsLetter bool
	}

// todo: slices
	var bookings = make([]UserDataStructure, 0)
	const conferenceTickets uint = 50
	var remainingTicket int= 50

// if the app is not wrapped inside for loop, use sync.wait
	// var wg = sync.WaitGroup{}




	func main(){
	 name  := "Go Conference"




	// greet

	greetUser(conferenceTickets, remainingTicket, name)

	// todo: creating arrays

	// var bookings [50]string
	// var booking []string --- slices



	for remainingTicket > 0 && len(bookings) < int(conferenceTickets) {
	
	isValidEmail, isValidName, firstName, lastName, email, userTickets :=helper.ValidateInput()

		fmt.Println(isValidName)

       //  todo: validate name
		if !isValidName {
			fmt.Printf("Your %v or %v is too short\n", firstName, lastName)
			continue
		}

    //   todo: validate email
		if  !isValidEmail {
           fmt.Printf("The provided email %v is not valid\n", email)
		   continue
		}


		// todo: check if userTicket is greater than remaining ticket


		if userTickets < 1{
			fmt.Printf("Our valued customer, you can't book %v tickets, please enter valid number\n", userTickets)
			break
		}
	
		if userTickets <= remainingTicket {
			// bookings[0] = firstName + " " + lastName
			// todo: using slice
	
				bookTicket( userTickets, firstName, lastName, name, email)
				// wg.Add(1)
				// todo adding GO infront of a function is called  multiThreading(avoids blocking) or new thread
				// todo: adding GO infront of function call generates a new thread to avoid code blocking while running go code
				// todo: when we create a thread, go creates GREEN Thread
				// todo: Creating hight level go routine
				// todo: Thread is cheaper to create, it's light weight, which means you can use thousands of thread without affecting the system
			    go sendTicket(userTickets, firstName, lastName,  email)
			 firstNames := getUserFirstName()

			fmt.Printf("The first names of all our bookings are : %v\n", firstNames)

			if remainingTicket == 0{
				// todo: quit the app
				fmt.Println("Our conference is booked-out.Come back next year")
				sendTicket(userTickets, firstName, lastName,  email)
				break
			}
        
		} else {

			fmt.Printf("Impossible, you can only book %v number of tickets\n", remainingTicket)
				
				// todo: continue specifies that the loop should go again if the first one fail
				// continue
		}
		//  wg.Wait()
	
	}

}



	func greetUser(conferenceTickets uint, remainingTicket int, name string){
		fmt.Printf("conferenceTicket is  %T, remainingTicket is %T.\n", conferenceTickets, remainingTicket)

		fmt.Printf("Welcome to %v booking application\n", name)
		fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTicket)
		fmt.Println("Get your tickets here to attend")
	}






	func bookTicket(  userTickets int, firstName string, lastName string, name string, email string) {


		// todo: create a map for a user. maps can only have same data-types for keys and same data types for values

		// var userData = make(map[string]string)
		// userData["firstName"] = firstName
		// userData["lastName"] = lastName
		// userData["email"] = email
		// userData["NumOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)


		var userData = UserDataStructure{
			firstName: firstName,
			lastName: lastName,
			email: email,
			isOptedInForNewsLetter: true,
			NumOfTickets: uint(userTickets),
		}


		remainingTicket = remainingTicket -userTickets
		bookings = append(bookings, userData)
		fmt.Printf("List of bookings are %v \n.", bookings)


		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n",firstName, lastName, userTickets, email)
		if userTickets > 0 {
			fmt.Printf("%v tickets remaining for %v \n", remainingTicket, name)
		}

	}


	func getUserFirstName() []string{
		firstNames :=[]string{}
			for _, booking := range bookings{
				// var names = strings.Fields(booking)
				
				// firstNames = append(firstNames, names[0])--- getting firstNames from sice
				// firstNames =append(firstNames, booking["firstName"]) --- getting firstNames from slices of map
				firstNames =append(firstNames, booking.firstName)
			}

			return firstNames
	}


	func sendTicket(userTickets int, firstName string, lastName string, email string){
		// todo sleep stops the thread from execusion for specified time limit, here, 10 second
		time.Sleep(10 * time.Second) 
		var ticket = fmt.Sprintf("%v tickets to %v %v", userTickets, firstName, lastName)
		fmt.Println("################################################")
		fmt.Printf("Sending : \n %v to email address %v.\n", ticket, email)
		fmt.Println("################################################")
        // wg.Done()
	}