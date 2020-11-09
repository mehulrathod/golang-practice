package main

import "fmt"

func main()  {
	sum := 0 //2
	for i := 1; i < 5; i ++ {
		if i % 2 != 0{ //More accurately anything that is not 0 is true. // skip odd numbers
			fmt.Println("-=-=-=value of i-=-=-=", i)
			continue // i = 2,4,6
		}
		fmt.Println("-=-=-=value of i-=-=-=", i)
		sum += i
		fmt.Println("-=-=-=value of sum-=-=-=", sum)
	}
	fmt.Println(sum) // 6(2+4)
}
