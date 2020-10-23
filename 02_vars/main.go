package main

import "fmt"

func main() {

	// main types:
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for uint32
	// float32 float64
	// complex64 complex128 (for really large numbers)

	// var name string = "Saad Ahmad"
	// technically do not need to add string type since it is inferred
	var fullName = "Saad Ahmad"

	// have to use vars or error
	var age = 22

	// bool
	var isCool = true

	// print
	fmt.Println("Full name and age:", fullName, age)

	// get type
	fmt.Printf("Type of isCool Bool: %T\n", isCool)

	// shorthand assignment (can only be used inside a func)
	firstName := "Saad"
	shoeSize := 11.5

	// shorthand
	lastName, email := "Ahmad", "ahmadsaad58"

	fmt.Println(firstName, shoeSize)

	fmt.Println(lastName, email)

}
