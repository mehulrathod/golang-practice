package main

import "fmt"

func main() {
	var p *int

	if p != nil {
		fmt.Println("P is having  a value: ", *p)
	} else {
		fmt.Println("Watchout for nil values")
	}

	var lifebooster float64 = 99.2
	lifeBoosterRef := &lifebooster //it holds the address/reference of pointer

	fmt.Println(lifebooster)
	fmt.Println(*lifeBoosterRef) //it holds the value of pointer
}
