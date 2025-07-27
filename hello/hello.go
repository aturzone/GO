package main

import (
	"fmt"
	"log"
	"github.com/aturzone/go/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//name input
	fmt.Print("Please enter your name: ")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal("Error reading input: ", err)
	}

	//greeting with name input
	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
