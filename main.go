package main

import (
	"fmt"
	"log"

	"example.com/go-htmx-app/greetings"
	"example.com/go-htmx-app/server"
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

	handler := server.RegisterEndpoints()
	server.StartServer(handler, "")
}
