package main

import (
	"fmt"
	"log"
	"github.com/aturzone/go/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Atur")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	message, err = greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
