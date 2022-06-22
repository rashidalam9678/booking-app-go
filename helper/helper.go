package helper
import "strings"


func ValidateUserInput(firstName string, lastName string, email string, numberOfTicketsToBeBooked int,remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := numberOfTicketsToBeBooked > 0 && numberOfTicketsToBeBooked <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}