package main

import (
	"booking_app/helper"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const tickets int = 50

var conferenceName = "MRA GO Conference"
var remainingTickets = 50
var UserResponse string = "y"

// var bookings = make([]map[string]string,1)

// struct
var bookings = []UserData{}

type UserData struct {
	firstName                 string
	lastName                  string
	email                     string
	numberOfTicketsToBeBooked int
}

func main() {
	greetUser()
	for UserResponse == "y" {
		firstName, lastName, email, numberOfTicketsToBeBooked := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, numberOfTicketsToBeBooked, remainingTickets)

		if isValidEmail && isValidName && isValidTickets {
			bookTicket(numberOfTicketsToBeBooked, firstName, lastName, email)
			sendCofirmationMail(firstName, lastName, email, numberOfTicketsToBeBooked)
		} else {
			if !isValidEmail {
				fmt.Println("You have entered the invalid email")
			}
			if !isValidName {
				fmt.Println("You have entered the invalid name")
			}
			if !isValidTickets {
				fmt.Println("You have entered the invalid number of tickets")
			}
		}

		if remainingTickets == 0 {
			fmt.Println("all the tickests have been sold")
			break
		}
		var userInput string
		fmt.Print("Do you want to book more tickets?(y/n) :")

		fmt.Scan(&userInput)
		UserResponse = userInput
	}

}

func greetUser() {
	fmt.Println("Welcome to", conferenceName, "Booking System")
	fmt.Println("Book your Tickets now, Only", remainingTickets, "are remaining")
}
func getBookings() []string {

	// for i := 0; i < len(bookings); i++ {} as standard for loop
	// for i<50 {} as standlard while loop
	// for{} as infinite loop
	// below is specially for arrays and slices

	bookingsFirstName := []string{}
	for _, booking := range bookings {

		bookingsFirstName = append(bookingsFirstName, booking.firstName)
	}
	return bookingsFirstName
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var numberOfTicketsToBeBooked int
	fmt.Println("Enter your First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to book")
	fmt.Scan(&numberOfTicketsToBeBooked)
	return firstName, lastName, email, numberOfTicketsToBeBooked
}

func bookTicket(numberOfTicketsToBeBooked int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - numberOfTicketsToBeBooked

	// map
	// var userData= make(map[string]string)
	// userData["firstName"]=firstName
	// userData["lastName"]= lastName
	// userData["email"]=email
	// userData["noOfTickets"]=strconv.FormatUint(uint64(numberOfTicketsToBeBooked),10)
	// bookings = append(bookings, userData)

	// struct method
	var userData = UserData{
		firstName:                 firstName,
		lastName:                  lastName,
		email:                     email,
		numberOfTicketsToBeBooked: numberOfTicketsToBeBooked,
	}
	bookings = append(bookings, userData)

	bookingsFirstName := getBookings()
	fmt.Println("Bookings are", bookingsFirstName)
	fmt.Println("All Users", bookings)

}
func sendCofirmationMail(firstName string, lastName string, email string, numberOfTicketsToBeBooked int) {

	
	time.Sleep(2 * time.Second)
	fmt.Println("You have booked", numberOfTicketsToBeBooked, "tickets for", firstName, lastName)
	fmt.Printf("##############################################################################\n")
	fmt.Println("Confirmation mail has been sent")
	fmt.Printf("##############################################################################\n")
}
