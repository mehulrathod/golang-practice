package main

import "fmt"

type Look interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length,width float64
}

func (r *Rectangle) Area() float64  {
	return r.width * r.length
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main()  {
	var s Look = &Rectangle{10.0, 5.0}
	fmt.Printf("Shape type = %T, Shape Value = %v %v \n", s,s,s)
	fmt.Printf("Area = %f, Perimeter = %f \n\n", s.Area(), s.Perimeter())
}
