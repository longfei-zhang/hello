package main

import (
	"fmt"
	"log"

	"github.com/longfei-zhang/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Girl", "Lily", "Saya"}

	//Request a greeting msg.
	//message, err := greetings.Hello("longfei")
	//Request greeting msg for the names.
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	//exit the program.
	if err != nil {
		log.Fatal(err)
	}

	//If no error was returned, print the returned message
	// to the console
	fmt.Println(messages)
}
