package main

import "fmt"

// Adder is a function that adds ints
func Adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	sum := Adder()

	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}

}
