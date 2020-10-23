package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// use * to read val from address
	fmt.Println(*b)  // 5
	fmt.Println(b)   // mem address
	fmt.Println(*&a) // 5

	// change val with pointer
	*b = 10
	fmt.Println(a)
}
