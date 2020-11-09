package main

import "fmt"

// struct is a collection of fields so we can group things together to create a more logical type
// we use struct to fulfil the needs of class and objects

type person struct {
	name string	// each fields need name and type
	age int
}

func main()  {
	p := person{name: "Mehul", age: 24}
	/*
	fmt.Println(p.age) //if u want particular field
	*/
	fmt.Println(p)
}