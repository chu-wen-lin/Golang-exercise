package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainTickets uint = 50
var bookings = make([]userData, 0) // slice size is dynamic, it will increase when we add new elements

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

// The main function is the starting point of a GO program
func main() {
	greetName()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)                                              // the number of threads that the main thread should wait for and
		go sendTicket(userTickets, firstName, lastName, email) // start a seperate thread for executing this logic to made our application concurrent with a go keyword
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	wg.Wait() // wait until the wait group counter is 0 and need to be put at the end of the main function
}

func greetName() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainTickets = remainTickets - userTickets

	// create a userData struct for a user
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second) // simulate that generating and sending tickets needs to take 50 seconds
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("-------------------")
	fmt.Printf("Sending tickets: \n %v \n to email address %v\n", ticket, email)
	fmt.Println("-------------------")
	wg.Done() // decrements the waiting group counter by 1
}
