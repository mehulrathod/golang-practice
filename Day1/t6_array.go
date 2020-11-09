package main

import "fmt"

func main()  {
	/*
	a := [5]int{1,1,2,12,4} //array
	*/
	a := []int{1,1,2,12,4} //slice
	a = append(a, 12)
	fmt.Println(a)
}
