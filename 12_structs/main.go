package main

import (
	"fmt"
	"strconv"
)

// Person is a struct
type Person struct {
	// firstName string
	// lastname  string
	// city      string
	// gender    string
	// age       int

	firstName, lastname, city, gender string
	age                               int
}

// Greet is a value receiver method
func (p Person) Greet() string {
	return "hello " + p.firstName + " " + p.lastname + "! You are " + strconv.Itoa(p.age) + " years old"
}

// HasBirthday is a pointer receiver method
func (p *Person) HasBirthday() {
	p.age++
}

// GetMarried is a pointer receiver method
func (p *Person) GetMarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastname = spouseLastName
	}
}

func main() {
	// initialize person
	person1 := Person{"saad", "ahmad", "austin", "m", 22}
	person2 := Person{firstName: "firstname", lastname: "lastname", city: "city", gender: "f", age: 0}

	// create people
	fmt.Println(person1, person2)
	// access attribute of person
	fmt.Println(person1.firstName)
	// modify attribute
	person2.age++
	fmt.Println(person2)
	// call pointer function
	person2.HasBirthday()
	// call value function
	fmt.Println(person1.Greet())

	person2.GetMarried("newlastname")
	fmt.Println(person2)

}
