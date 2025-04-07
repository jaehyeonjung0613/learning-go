package main

import (
	"fmt"
	"log"

	"example/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Ldate)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
