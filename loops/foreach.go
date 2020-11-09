package main

import "fmt"

/*Looping over elements in slices, arrays, maps, channels or strings is often better done with a range loop.*/

func main()  {
	//strings := []string{"hello", "world"}
	integer := []int64{1,2,4,5,6}

	for i,s := range integer {
		fmt.Println("-=-=Value of i-==-=-=",i)
		fmt.Println("-=-=Value of s-==-=-=",s)
		fmt.Println(i,s)
	}
}
