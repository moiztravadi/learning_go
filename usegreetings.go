package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	// Get a greeting message and print it.
	// message, err := greetings.Hello("Moiz")

	// get multiple greetings
	names := []string{"Moiz", "Ibrahim"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for key, message := range messages {
		fmt.Printf("For key %s, Value is %s \n", key, message)
	}

}
