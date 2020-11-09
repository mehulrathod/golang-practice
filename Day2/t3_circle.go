package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
	Diameter() float64
}

//type Rectangle struct {
//	length,width float64
//}

type Circle struct {
	Radius float64
}

const pi float64 = 3.14 //declaration of constant
/*func (r * Rectangle) Area() float64  {
	return r.width * r.length
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}*/

func (c *Circle) Area() float64  {
	return pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * pi * c.Radius
}

func (c *Circle) Diameter() float64  {
	return 2 * c.Radius
}

func main()  {
	Shape := Circle{5.0}
	fmt.Printf("Shape type = %T, Shape Value = %v \n", Shape,Shape)
	fmt.Printf("Area = %f, Perimeter = %f \n\n", Shape.Area(), Shape.Perimeter())

	//Shape := Rectangle{10.0,5.0}
	//fmt.Printf("Shape type = %T, Shape Value = %v %v \n", Shape,Shape,Shape)
	//fmt.Printf("Area = %f, Perimeter = %f \n\n", Shape.Area(), Shape.Perimeter())
}