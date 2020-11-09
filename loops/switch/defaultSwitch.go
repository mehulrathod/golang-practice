package main

import (
	"fmt"
	"time"
)

func main() {
/*	m := time.Monday
	switch time.Now().Weekday() {
	case m, time.Saturday:
		fmt.Printf("Today is Saturday. %T \n", time.Now())
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Today is a weekday.")
	}*/
	switch hour := time.Now().Hour(); { // missing expression means "true"
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
