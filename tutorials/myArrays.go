package main

import "fmt"

func main() {
	var numbers [3]string
	numbers[0] = "uno"
	numbers[1] = "aste"
	numbers[2] = "karate"
	fmt.Println(numbers)

	colors := [4]string{"ram", "laxman", "sita", "parvati"}
	fmt.Println(colors)
	fmt.Println(colors[2])
	fmt.Println(len(colors))
}
