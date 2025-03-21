package main

import (
	"fmt"
	"log"

	"example.com/go-htmx-app/greetings"
	"rsc.io/quote"
)

func main() {
	test := 12345
	fmt.Println(test)
	fmt.Println(quote.Glass())
	greeting, err := greetings.Hello("Alex")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(greeting)
	}
}
