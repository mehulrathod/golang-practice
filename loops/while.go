package main

import "fmt"

func main()  {
	i := 1 // i = 1,2,4
	for i < 5 {
		fmt.Println("-=-=-IIII-=-=-=", i)
		i *= 2 // i * 1,2,4
	}
	fmt.Println(i) // 8
}
