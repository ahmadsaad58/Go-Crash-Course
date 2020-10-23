package main

import "fmt"

func main() {
	// make map
	emails := make(map[string]string)

	// assign values
	emails["bob"] = "bob@gmail.com"
	emails["bob2"] = "bob2@gmail.com"
	emails["bob3"] = "bob3@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["bob"])

	// delete
	delete(emails, "bob3")
	fmt.Println(emails)

	// explicit
	newEmails := map[string]string{"bob": "bob@gmail.com"}
	fmt.Println(newEmails)

}
