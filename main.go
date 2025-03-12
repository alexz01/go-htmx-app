package main

import (
	"fmt"

	"example.com/go-htmx-app/greetings"
	"rsc.io/quote"
)

func main() {
	test := 12345
	fmt.Println(test)
	fmt.Println(quote.Glass())
	fmt.Println(greetings.Hello("Alex"))
}
