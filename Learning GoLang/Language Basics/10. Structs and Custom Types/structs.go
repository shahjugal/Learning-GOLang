package main

import (
	"fmt"

	"example.com/structs/person"
)

// We can alias also like typedef of c++
type str string

// and can also attach the methods to aliased class which wont change the original one.

func (s *str) log() {
	fmt.Println(*s)
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var p1 *person.Person = person.New(
		firstName,
		lastName,
		birthdate,
	)

	// ... do something awesome with that gathered data!
	p1.PrintPerson()
	p1.ClearPerson()
	p1.PrintPerson()

	email := getUserData("Please enter your email: ")
	password := getUserData("Please enter your password: ")

	var p2 *person.Admin = person.NewAdmin(
		email,
		password,
	)

	// ... do something awesome with that gathered data!
	p2.PrintPerson()
	p2.ClearPerson()
	p2.PrintPerson()

	var msg str = "Ending the program now..."
	msg.log()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
