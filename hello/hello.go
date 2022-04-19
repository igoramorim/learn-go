package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(1) // Shows the current date when logging

	message, err := greetings.Hello("Igor")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string {"Igor", "Maria", "João"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}