package main

import "fmt"

func main() {
	// arrays
	var fruitArray [2]string

	// assign vals
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	fmt.Println(fruitArray)

	// decare and assign
	shortFruitArray := [2]string{"Apple", "Orange"}

	fmt.Println(shortFruitArray)

	// slices
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Printf("The type of fruit slice: %T\nThe type of fruit array: %T\n", fruitSlice, fruitArray)
	// The type of fruit slice: []string
	// The type of fruit array: [2]string

	// len
	fmt.Println(len(fruitSlice))

	// range
	fmt.Println(fruitSlice[1:2])

}
