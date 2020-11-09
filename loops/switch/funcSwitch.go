package main

import "fmt"

// Foo prints and returns n.
func Foo(n int) int {
	fmt.Println(n)
	return n
}

func main() {
	switch Foo(2) {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("First case")
		fallthrough
	case Foo(4):
		fmt.Println("Second case")
	}
}
