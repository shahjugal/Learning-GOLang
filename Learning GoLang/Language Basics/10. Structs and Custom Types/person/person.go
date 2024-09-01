package person

import (
	"fmt"
	"time"
)

// Creating struct
type Person struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Creating struct of struct
type Admin struct {
	email    string
	password string
	// person   Person	// this works well
	Person // but this is call anonymous embedding which allows call Admin.ClearPerson instead of Admin.person.ClearPerson
}

// this method is now attached to the struct person
func (p Person) PrintPerson() {
	fmt.Println(p.firstName, p.lastName, p.birthdate)

}

// this is constructor method for person
func New(firstName string, lastName string, birthdate string) *Person {
	return &Person{firstName, lastName, birthdate, time.Now()}

}

func NewAdmin(email string, password string) *Admin {
	return &Admin{email: email, password: password, Person: Person{"ADMIN", "ADMIN", "", time.Now()}}
}

// this method is now attached to the struct person
func (p *Person) ClearPerson() {
	p.firstName = ""
	p.lastName = ""

}
