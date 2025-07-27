package main

import (
	"fmt"
	"github.com/aturzone/go/greetings"
)

func main() {
	message := greetings.Hello("Atur")
	fmt.Println(message)
}
