package main

import (
	"fmt"
	"log"
	"github.com/aturzone/go/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greeting.Hello("Atur")
	if err != nill {
		log.Fatal(err)
	}
	fmt.Println(message)

	message, err = greeting.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
