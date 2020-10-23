package main

import "fmt"

// Greeting is a function that returns Hello {Name}
func Greeting(name string) string {
	return "Hello " + name
}

// GetSum is a function that returns the sum of 2 numbers
func GetSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(Greeting("Saad"))
	fmt.Println(GetSum(1, 3))
}
