package main

import "fmt"

func main() {
	score := make(map[string]int)
	score["Mehul"] = 89
	fmt.Println(score)

	score["Hitesh"] = 89
	score["Ramesh"] = 99
	score["Rakesh"] = 45
	score["Suresh"] = 72
	fmt.Println(score)

	getRakeshScore := score["Rakesh"]
	fmt.Println(getRakeshScore)

	delete(score, "Suresh")
	fmt.Println(score)

	for k, v := range score {
		fmt.Printf("Score of %v is %v\n", k, v)
	}
}
