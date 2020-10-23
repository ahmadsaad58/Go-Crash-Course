package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>hello world</h1>")
}

func hello(w http.ResponseWriter, r *http.Request) {
	names, avail := r.URL.Query()["name"]

	if !avail || len(names) < 1 {
		fmt.Fprint(w, "ah shoot, you forgot the name")
		return
	}

	name := names[0]
	fmt.Fprint(w, "hello "+name)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	fmt.Println("Server Started")
	http.ListenAndServe(":3000", nil)
}
