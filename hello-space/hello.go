package main

import (
	"fmt"
	"log"

	"github.com/ravidhavlesha/learn-go/greetings"
)

func main() {
	log.SetPrefix(("Greetings:"))
	log.SetFlags(0)

	fmt.Println("Hello, Space!")
	message, err := greetings.Greet("Ravi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Ravi", "Shakti", "Dharm"}
	messages, err := greetings.Greets(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
