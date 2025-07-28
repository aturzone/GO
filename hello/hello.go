package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
	"example.com/utils"
	)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(24601), reverse.ReverseSlice([]string{"a","b","c"}), utils.SayHi())
}
