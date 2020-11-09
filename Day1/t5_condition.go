package main

import "fmt"

func main()  {
	x := -9
	if x > 6 {
		fmt.Println(x, "%s is greater than 6")
	} else if x == 6 {
		fmt.Printf("5 is equals to 6")
	} else if x <= 6 {
		fmt.Println("5 is lesser than 6")
	} else {
		fmt.Println("%s is better", x )
	}
}
