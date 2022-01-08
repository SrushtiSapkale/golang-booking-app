package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var confname string = "GO BOOK TICKETS"
var conftickets int = 50
var remainingtickets int = 50
var bookings = make([]Userdata, 0)

type Userdata struct {
	username  string
	email     string
	notickets int
}

var wg = sync.WaitGroup{}

func main() {
	greet()

	//	for remainingtickets > 0 && len(bookings) < 50 {
	//user i/p
	username, email, usertickets := userip()
	//validate i/p
	isvalidname, isvalidemail, isvalidticketno := validateinput(username, email, usertickets, remainingtickets)
	//if i/p is valid
	if isvalidname && isvalidemail && isvalidticketno {
		if usertickets < remainingtickets {
			booktickets(email, username, usertickets)
			//send ticket will become another thread. meanwhile the next command is executed
			//and when the  sendticket func gets completed it will get executed in middle without interrupting current pgm

			//generating and sending ticket task will now run in the background
			wg.Add(1)
			go sendticket(usertickets, username, email)

			if remainingtickets == 0 {
				fmt.Printf("\nTICKETS BOOKED OUT!")
				//break
			}
		} else {

			fmt.Printf("Sorry, we only have %v tickets\n", remainingtickets)

		}

	} else {
		if !isvalidname {
			fmt.Printf("\nName should be more than two characters.")
		}
		if !isvalidemail {
			fmt.Printf("\nEmail should be a valid one please! ")
		}
		if !isvalidticketno {
			fmt.Printf("\nTicket number should be less than %v\n", remainingtickets)
		}
	}

	wg.Wait()
}

func greet() {
	fmt.Println("WELCOME")
	fmt.Println("Get your tickets booked at", confname)
	fmt.Println("we have", conftickets, "tickets and only ", remainingtickets, "are remaining")
}

func validateinput(username string, email string, usertickets int, remainingtickets int) (bool, bool, bool) {
	var isvalidname = len(username) >= 2
	var isvalidemail = strings.Contains(email, "@")
	isvalidticketno := usertickets > 0 && usertickets <= remainingtickets

	return isvalidname, isvalidemail, isvalidticketno
}

func userip() (string, string, int) {

	var username string
	var email string
	var usertickets int
	fmt.Println("Enter username :")
	fmt.Scan(&username)

	fmt.Println("Enter email :")
	fmt.Scan(&email)

	fmt.Println("Enter no. of tickets :")
	fmt.Scan(&usertickets)

	return username, email, usertickets
}

func booktickets(email string, username string, usertickets int) {
	fmt.Printf("\nThankyou so much %v for booking %v tickets with us.\nYou will soon get a confirmation email at %v. \n", username, usertickets, email)

	remainingtickets = remainingtickets - usertickets
	fmt.Printf("\nThere are now %v tickets remaining.", remainingtickets)

	//user map
	var userdata = Userdata{
		username:  username,
		email:     email,
		notickets: usertickets,
	}

	bookings = append(bookings, userdata)
	fmt.Printf("\nThese are our bookings: %v\n", bookings)

}

func sendticket(usertickets int, username string, email string) {
	time.Sleep(50 * time.Second)
	var tick = fmt.Sprintf("\n'%v you have %v tickets", username, usertickets)
	fmt.Printf("##################")
	fmt.Printf("\nSending ticket %v'\nto email address %v\n", tick, email)
	fmt.Printf("##################\n")
	wg.Done()

}
