package main

import (
	"fmt"
	"log"

	"example/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Ldate)

	message, err := greetings.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
