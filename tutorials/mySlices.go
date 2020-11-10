package main

import (
	"fmt"
	"sort"
)

func main() {
	var pets = []string{"cat", "dog", "horse"}
	fmt.Println(pets)

	pets = append(pets, "rabbits")
	fmt.Println(pets)

	pets = append(pets[1:])
	fmt.Println(pets)

	pets = append(pets[1 : len(pets)-1])
	fmt.Println(pets)

	var heros = make([]string, 3, 3)
	heros[0] = "thor"
	heros[1] = "spiderman"
	heros[2] = "ironman"
	fmt.Println(heros)

	heros = append(heros, "deadpool")
	fmt.Println(heros)

	myThings := []int{754, 27, 87, 47, 2, 41, 45}
	sort.Ints(myThings)
	fmt.Println(myThings)
}
