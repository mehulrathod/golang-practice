package main

import (
	"errors"
	"fmt"
	"math"
)

func calc(x int, y int) int  {
	return x+y
}

func divide(n int, m int) int {
	if n < m{
		fmt.Println("its not valid, n have to greater than m")
	}
	return n/m
}

func sqrt(n float64) (float64, error)  {
	if n < 0{
		return 0, errors.New("undefined for negative number")
	}

	return math.Sqrt(n), nil
}

func main()  {
	// for addition
	result := calc(1,2)
	fmt.Println(result)

	// for division
	dm := divide(50,5)
	fmt.Println(dm)

	// sqrt
	res,err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
