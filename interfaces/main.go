package main

import "fmt"

type Shape interface {
	Area() float64
	Diameter() float64
	Circumference() float64
}

type Circle struct {
	radius float64
}

const pi = 3.14

func (c *Circle) Area() float64 {
	return pi * c.radius * c.radius
}

func (c *Circle) Diameter() float64 {
	return 2 * c.radius
}

func (c *Circle) Circumference() float64  {
	return 2 * pi * c.radius
}

func main()  {
	c := Circle{0.5}
	var s Shape = &c
	area := s.Area()
	diameter := s.Diameter()
	circumference := s.Circumference()
	fmt.Println("Area of Circle is : ",area)
	fmt.Println("Area of Diameter is : ",diameter)
	fmt.Println("Area of Circumference is : ",circumference)
/*	fmt.Printf("Shape type = %T, Shape Value = %v \n", s,s)
	fmt.Printf("Area = %f, Perimeter = %f \n\n", s.Area(), s.Circumference())
	fmt.Printf("Area = %f, Perimeter = %f \n\n", s.Area(), s.Circumference())
*/
	fmt.Printf("Area = %f, Circumference = %f \n\n", s.Area(), s.Circumference())
	fmt.Printf("Area = %f, Diameter = %f \n\n", s.Area(), s.Diameter())
}
