package main

import "fmt"

// function which swap values
func swap(a, b *int)int{
	var o int
	o = *a
	*a = *b
	*b = o

	return o
}

// Main function
func main() {

	var p int = 10
	var q int = 20
	fmt.Printf("p = %d and q = %d", p, q)

	// call by reference
	swap(&p, &q)
	fmt.Printf("\np = %d and q = %d",p, q)
}

/*NOTE : Call by reference: Both the actual and formal parameters refer to the same locations, so any changes made
inside the function are actually reflected in actual parameters of the caller.*/
