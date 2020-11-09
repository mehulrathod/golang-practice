package main

import "fmt"

// function which swap values
func swapp(a, b int)int{

	var o int
	o= a
	a=b
	b=o

	return o
}

// Main function
func main() {
	var p int = 10
	var q int = 20
	fmt.Printf("p = %d and q = %d", p, q)

	// call by values
	swapp(p, q)
	fmt.Printf("\np = %d and q = %d",p, q)
}

/*NOTE: Call by value: : In this parameter passing method, values of actual parameters are copied to function’s formal
parameters and the two types of parameters are stored in different memory locations. So any changes made inside
functions are not reflected in actual parameters of the caller.*/