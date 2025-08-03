package main

import (
	"fmt"
	"time"
)

func main() {
	live := time.Now()
	today := live.Weekday()
	hour := live.Hour()

	switch {
	case hour < 12:
		fmt.Println("Hi! Good Morning.")
	case hour < 18:
		fmt.Println("Hi! Good afternoon.")
	default:
		fmt.Println("Hi! Good evening.")
	}

	switch today {
	case time.Saturday:
		fmt.Println("We are in Saturday.")
	case time.Sunday:
		fmt.Println("We are in Sunday.")
	case time.Monday:
		fmt.Println("We are in Monday.")
	case time.Tuesday:
		fmt.Println("We are in Tuesday.")
	case time.Wednesday:
		fmt.Println("We are in Wednesday.")
	case time.Thursday:
		fmt.Println("We are in Thursday.")
	case time.Friday:
		fmt.Println("We are in Friday.")
	default:
		fmt.Println("Unknown day.")
	}
}
