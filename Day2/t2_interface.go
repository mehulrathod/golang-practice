package main

import (
	"fmt"

)

type Actions interface {
	Walk(distance float64)
	Run(distance float64)
}

//type Person struct {
//	Name string
//	Age int64
//}
type Animal struct {
	Name, Species string
}

//func (p *Person) Walk(distance float64) (n int, err error) {
//	return fmt.Printf("%s walks : %g kms \n", p.Name, distance)
//}

func (a *Animal) Walk(distance float64) (int, error) {
	return fmt.Printf("%s %s walks : %g kms \n", a.Name, a.Species, distance)
}
func (a *Animal) Run(distance float64) (int, error) {
	return fmt.Printf("%s %s runs : %g kms \n", a.Name, a.Species, distance)
}

func main()  {
	//person := Person{"Mehul", 20}
	//person.Walk(20 )

	animal := &Animal{"Dog", "German"}
	animal.Walk(20 )
	animal.Run(4.5 )
}
