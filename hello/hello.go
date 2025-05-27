package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request a greeting message fro a single person
	// message, err := greetings.Greet("Gladys")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// request a greeting message for multiple people
	// as slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// request greeting messages for the names
	messages, err := greetings.Greets(names)
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the message to the console
	fmt.Println(messages)
}
