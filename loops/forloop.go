package main

import "fmt"

func main()  {
	for i := 0; i <= 4; i++ {
		for j := 0; j <= i; j++{
			fmt.Printf("*")
		}
		fmt.Printf(" \n")
	}
}