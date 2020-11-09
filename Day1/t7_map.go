package main

import "fmt"

func main()  {
/*	vertices := make(map[int]int)
	vertices[1] = 5
	vertices[2] = 15
	vertices[3] = 25
	vertices[4] = 35*/
	vertices := make(map[string]int)
	vertices["apple"] = 5
	vertices["chilly"] = 15
	vertices["beetroot"] = 25
	vertices["pineapple"] = 35

	delete(vertices, "beetroot")

	fmt.Println(vertices)
}
