package main

import "fmt"

type Person struct {
	Name string
	Age int64
}

func (p *Person) Talk(pharse string) (int, error) {
	return fmt.Printf("%s says : %s\n", p.Name, pharse)
}

func main() {
	person := &Person{"Mehul", 24}
	person.Talk("Go is awesome")
}