package main

import "fmt"

func main() {
	ids := []int{1, 2, 3}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// sum items
	sum := 0
	for _, val := range ids {
		sum += val
	}
	fmt.Println(sum)

	// range map
	emails := make(map[string]string)
	emails["bob"] = "bob@gmail.com"
	emails["bob2"] = "bob2@gmail.com"
	emails["bob3"] = "bob3@gmail.com"

	for key, val := range emails {
		fmt.Printf("%s: %s\n", key, val)
	}

}
